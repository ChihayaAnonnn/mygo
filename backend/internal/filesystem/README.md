# Filesystem Storage Layer（对象存储式）设计规范

> 本文档定义一个**更底层的 filesystem/storage 层**：以 `key -> bytes` 的对象存储抽象为核心。
> 该层不理解任何上层领域概念（如 Knowledge、agent、public/private 等）；只提供稳定的存取语义与实现约束。
> 上层系统通过“key schema（命名规则）”把业务对象映射到存储 key。

## 1. 目标与边界

### 1.1 目标

- **领域无关**：不依赖上层业务类型与命名，不把业务字段写入路径规范。
- **稳定存取语义**：`Put/Get/Delete/Head/List` 等操作具备明确一致性/覆盖规则。
- **写入安全**：提供可实现“原子替换（atomic replace）”的写入语义，避免半写入。
- **可替换后端**：同一抽象可由 Local FS、S3/MinIO、NFS 等后端实现。

### 1.2 非目标

- **不定义鉴权模型**：访问控制由上层完成；storage 层只接收已授权的请求上下文（若需要）。
- **不定义业务版本控制**：业务版本（如文档版本）属于上层事实源（通常是 DB）；storage 的“版本化”仅作为可选能力。
- **不保证人类手改稳定**：storage 输出的对象可能被上层重新生成并覆盖；若要保留手改，应在上层引入专门的 import/merge 流程。

## 2. 核心概念

- **basePath**：存储根目录（Local FS 后端）或 bucket/prefix（对象存储后端）。部署默认值可以是 `/workspace/data/objects`，但这只是路径选择，不携带领域语义。
- **namespace**：逻辑隔离域，用于区分“不同库/不同租户/不同用途”的对象集合。namespace 为不透明字符串，由上层决定。
  - 约束：只能包含安全字符集（建议：`[a-z0-9][a-z0-9._-]{0,63}`），避免路径穿越。
  - 说明：namespace 的意义由上层解释，storage 层不内置 `public/agents/...` 等含义。
- **key**：对象键（相对路径风格的字符串），由上层通过 key schema 生成。
  - 建议：使用分段（`/`）组织前缀，便于 List/GC。
  - 约束：禁止 `..`、绝对路径、反斜杠等，避免路径穿越。
- **object**：存储对象。包含 `bytes` 与可选 `metadata`（键值对）、`contentType`、`etag`/`checksum`、`lastModified` 等。

## 3. 接口（语义层）

### 3.1 基础操作

- `Put(namespace, key, bytes, opts)`：写入对象。
  - 默认语义：覆盖同名 key。
  - 可选：`IfMatch(etag)` / `IfNoneMatch`（CAS）用于并发保护（实现可选）。
- `Get(namespace, key)`：读取对象（返回 bytes + metadata）。
- `Head(namespace, key)`：仅取元信息（大小、etag、lastModified、metadata）。
- `Delete(namespace, key)`：删除对象（幂等：不存在也视为成功）。
  - 默认语义（本阶段）：**软删除**（tombstone）。在元数据存储中标记 `deleted_at`。
  - 约束（本阶段）：**不物理删除 payload**，以保证对象可恢复。
- `List(namespace, prefix, limit, continuationToken)`：按前缀列出 key（用于索引、GC、同步）。
  - 默认：不返回已软删除对象（`deleted_at IS NULL`）。
  - 可选：支持 `includeDeleted=true` 用于审计/对账/回收。

#### 软删除与恢复（本阶段）

若底层实现采用数据库表维护对象元信息（例如 `fs_objects`），建议增加两条约定：

- `Undelete(namespace, key)`（可选）：把 `deleted_at` 清空以恢复对象（仅当 payload 仍在且未被回收时可用）。
  - 本阶段由于不做 GC/回收，payload 默认仍在，通常可直接恢复。

### 3.2 原子替换（推荐支持）

提供“原子替换”能力的两种实现方式之一：

- **API 语义**：`PutAtomic(namespace, key, bytes, opts)`，承诺写入对读者呈现为一次性替换。

或

- **实现约束**：对 Local FS 后端，`Put` 采用临时文件 + `rename` 的实现策略（见 4.2）。

说明：对 S3 这类后端，单对象 `PUT` 本身就是原子可见（覆盖时读者看到旧或新），但需定义一致性与缓存策略。

## 4. Local FS 后端实现约束（参考实现）

### 4.1 路径映射

在 Local FS 后端，物理路径映射建议为：

```text
<basePath>/namespaces/<namespace>/<key>
```

理由：

- `namespaces/` 明确隔离层级，便于运维与清理。
- namespace 与 key 都由上层决定，但必须通过校验/编码避免路径穿越。

### 4.2 原子写入策略

对同一个 `(namespace, key)`，写入必须满足“要么完全是旧内容，要么完全是新内容”：

1. 确保父目录存在：`mkdir -p <dir>`。
2. 写临时文件：`<key>.tmp`（建议随机后缀，避免并发覆盖）。
3. `fsync(tmp)`（可选：性能/可靠性权衡）。
4. `rename(tmp, target)` 原子替换。
5. （可选）`fsync(dir)`，提高崩溃一致性。

### 4.3 并发与覆盖规则

- 默认：**最后写入者获胜**。
- 若上层需要强一致，应使用 `IfMatch(etag)` 语义或在上层做互斥（锁/事务）。

## 5. key schema（上层策略，不属于 storage 规范）

storage 层不定义业务 key schema，但建议 key 满足：

- **可分段**：便于按前缀 list，例如 `objects/<type>/<id>/latest`。
- **可 GC**：能通过前缀枚举并清理过期对象。
- **可调试**：必要时在 key 中放入短可读前缀，但不要把敏感信息放进 key。

### 5.1 示例：把“某个库的 Markdown 投影”映射为对象 key（仅示例）

> 这是一个上层（例如 Knowledge 投影服务）可能采用的 schema，用于说明如何使用 namespace/key。
> 注意：示例不构成 storage 规范的一部分。

- namespace：`<space_id>`（上层称为库 ID/space ID，作为不透明字符串）
- key：
  - `nodes/<node_id>/latest.md`
  - `nodes/<node_id>/assets/<filename>`

物理路径（Local FS）：

```text
<basePath>/namespaces/<space_id>/nodes/<node_id>/latest.md
```

## 6. 数据库持久化模型（`fs_objects`）

元数据通过 `ObjectPO`（Persistence Object）存入 `fs_objects` 表。表只存元数据，**不存 payload（bytes）**；同一 `(namespace, key)` 只保留一条记录，不做历史版本化。

### 6.1 字段说明

#### 标识与寻址

| 字段 | 类型 | 说明 |
|------|------|------|
| `id` | `bigint` AUTO_INCREMENT | 数据库内部主键，业务层不感知。 |
| `namespace` | `varchar(64)` NOT NULL | 逻辑隔离域。参与联合唯一索引 `uk_fs_objects_namespace_key`，并有单独索引 `idx_fs_objects_namespace` 用于按命名空间批量查询。具体含义由上层解释，storage 层不内置任何业务语义。 |
| `key` | `text` NOT NULL | 对象键（相对路径风格字符串），由上层通过 key schema 生成。与 `namespace` 共同构成联合唯一约束，保证同一命名空间内 key 唯一。 |

#### 后端与内容

| 字段 | 类型 | 说明 |
|------|------|------|
| `backend` | `varchar(16)` NOT NULL DEFAULT `'local'` | 存储后端标识，声明 payload 实际存储在哪种后端（`local` / `s3` / `minio` 等）。元数据表保持后端无关，只记录路由指针。 |
| `content_type` | `varchar(128)` | MIME 类型（如 `text/markdown`、`image/png`）。供 `Get` 时返回正确的 HTTP `Content-Type`，也可用于按类型过滤。 |
| `size_bytes` | `bigint` NOT NULL DEFAULT `0` | payload 字节大小。`Head` 操作无需读取文件即可返回大小，也用于配额统计。 |

#### 完整性校验

| 字段 | 类型 | 说明 |
|------|------|------|
| `etag` | `varchar(128)` | 实体标签（通常为内容哈希的十六进制或 quoted-string 格式）。用于 `IfMatch` / `IfNoneMatch` 条件请求，实现 CAS（比较并交换）并发保护。 |
| `checksum` | `varchar(128)` | 内容校验和（如 `sha256:<hex>`）。语义比 ETag 更严格：用于数据完整性验证，防止 payload 损坏或传输错误。 |

#### 扩展元数据

| 字段 | 类型 | 说明 |
|------|------|------|
| `metadata` | `jsonb` | 任意键值对，供上层附加业务标注（如 `source_version`、`encoding`、自定义标签等）。storage 层不解释其内容。 |

#### 时间戳

| 字段 | 类型 | 说明 |
|------|------|------|
| `created_at` | `timestamptz` NOT NULL | GORM 自动维护的创建时间。 |
| `updated_at` | `timestamptz` NOT NULL | GORM 自动维护的最后更新时间。 |
| `deleted_at` | `timestamptz` | GORM 软删除字段。非 NULL 时对象被视为已删除，`List` 默认过滤（`deleted_at IS NULL`）；物理 payload 保留，对应设计规范中的"软删除（tombstone）"语义。有索引 `idx_fs_objects_deleted_at`。 |

### 6.2 索引汇总

| 索引名 | 字段 | 类型 | 说明 |
|--------|------|------|------|
| `pk` | `id` | PRIMARY KEY | 主键 |
| `uk_fs_objects_namespace_key` | `(namespace, key)` | UNIQUE | 保证同一命名空间内 key 唯一 |
| `idx_fs_objects_namespace` | `namespace` | INDEX | 按命名空间批量查询（List、GC） |
| `idx_fs_objects_deleted_at` | `deleted_at` | INDEX | 软删除过滤加速 |

## 7. 扩展能力（可选）

- **对象版本化**：由后端能力提供（如 S3 Versioning），或在 key schema 中显式引入版本段。
- **元数据索引**：若需要按 metadata 检索，应由上层 DB/索引系统承担。
- **垃圾回收（GC）**：暂不实现。若未来需要物理清理软删除对象，再补充回收接口与保留策略。
