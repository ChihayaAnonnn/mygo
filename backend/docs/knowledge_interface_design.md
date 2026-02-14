# Knowledge 接口设计

> mygo 定位为 Data Plane Service，本文档定义 Knowledge 模块的 Service 接口和 HTTP API。
>
> 设计目标：
>
> - 接口即架构（Interface = Architecture）
> - mygo 只负责数据 CRUD + 向量检索

---

## 1. KnowledgeService（知识元服务）

**职责**：

- 管理 Knowledge Node 的元信息与生命周期
- 不关心具体内容，只关心元信息

```go
type KnowledgeService interface {
    CreateKnowledge(ctx context.Context, cmd CreateKnowledgeCmd) (KnowledgeID, error)
    UpdateKnowledgeMeta(ctx context.Context, cmd UpdateKnowledgeMetaCmd) error
    GetKnowledge(ctx context.Context, id KnowledgeID) (*Node, error)
    ListKnowledge(ctx context.Context, query KnowledgeQuery) ([]*Node, error)
    ArchiveKnowledge(ctx context.Context, id KnowledgeID) error
}
```

---

## 2. KnowledgeVersionService（版本服务）

**职责**：

- 管理 Markdown 内容的版本演化
- `knowledge_versions` 是事实源

```go
type KnowledgeVersionService interface {
    CreateVersion(ctx context.Context, cmd CreateVersionCmd) (VersionID, error)
    GetVersion(ctx context.Context, id VersionID) (*Version, error)
    GetLatestVersion(ctx context.Context, knowledgeID KnowledgeID) (*Version, error)
    ListVersions(ctx context.Context, knowledgeID KnowledgeID) ([]*Version, error)
}
```

---

## 3. MarkdownRenderService（文件派生服务，预留）

**职责**：

- 将 DB 中的 Markdown 内容派生为文件系统产物
- FS 永远不是事实源

```go
type MarkdownRenderService interface {
    RenderToFile(ctx context.Context, version *Version) (FilePath, error)
    RemoveFile(ctx context.Context, versionID VersionID) error
}
```

---

## 4. KnowledgeApplicationService（应用层编排）

**职责**：

- 编排多个底层 Service
- 对外提供高层用例

```go
type KnowledgeApplicationService interface {
    PublishKnowledge(ctx context.Context, knowledgeID KnowledgeID) error
    RebuildIndex(ctx context.Context, knowledgeID KnowledgeID) error
}
```

**说明**：

- `PublishKnowledge`：更新状态为 published + 渲染文件（如有）
- `RebuildIndex`：清除旧的 Chunk/Embedding 数据，Agent 端负责重新生成并写入

---

## 5. EpisodeService（数据摄入事件服务）

**职责**：

- 管理数据摄入事件（Episode）的生命周期
- 提供数据溯源能力

```go
type EpisodeService interface {
    CreateEpisode(ctx context.Context, cmd CreateEpisodeCmd) (EpisodeID, error)
    GetEpisode(ctx context.Context, id EpisodeID) (*Episode, error)
    ListEpisodes(ctx context.Context, query EpisodeQuery) ([]*Episode, error)
    AddMentions(ctx context.Context, episodeID EpisodeID, entityIDs []EntityID) error
    ListMentionedEntities(ctx context.Context, episodeID EpisodeID) ([]*Entity, error)
}
```

---

## 6. EntityService（实体服务）

**职责**：

- 管理从文档/对话中抽取的细粒度实体
- 提供实体的语义搜索能力

```go
type EntityService interface {
    CreateEntity(ctx context.Context, cmd CreateEntityCmd) (EntityID, error)
    BatchCreateEntities(ctx context.Context, cmds []CreateEntityCmd) ([]EntityID, error)
    GetEntity(ctx context.Context, id EntityID) (*Entity, error)
    UpdateEntity(ctx context.Context, cmd UpdateEntityCmd) error
    DeleteEntity(ctx context.Context, id EntityID) error
    ListEntities(ctx context.Context, query EntityQuery) ([]*Entity, error)
    SearchEntities(ctx context.Context, embedding EmbeddingVector, topK int) ([]*Entity, error)
}
```

---

## 7. EntityEdgeService（实体关系服务）

**职责**：

- 管理实体间的三元组关系
- 支持双时态查询和边失效

```go
type EntityEdgeService interface {
    CreateEntityEdge(ctx context.Context, cmd CreateEntityEdgeCmd) (EntityEdgeID, error)
    BatchCreateEntityEdges(ctx context.Context, cmds []CreateEntityEdgeCmd) ([]EntityEdgeID, error)
    GetEntityEdge(ctx context.Context, id EntityEdgeID) (*EntityEdge, error)
    InvalidateEntityEdge(ctx context.Context, id EntityEdgeID) error
    ListEdgesByEntity(ctx context.Context, entityID EntityID) ([]*EntityEdge, error)
    SearchEntityEdges(ctx context.Context, embedding EmbeddingVector, topK int) ([]*EntityEdge, error)
}
```

---

## 8. CommunityService（社区服务）

**职责**：

- 存储和查询社区检测结果
- 社区检测算法由 ave_mujica 执行

```go
type CommunityService interface {
    RebuildCommunities(ctx context.Context, communities []CreateCommunityCmd) error
    GetCommunity(ctx context.Context, id CommunityID) (*Community, error)
    ListCommunities(ctx context.Context, query CommunityQuery) ([]*Community, error)
    ListCommunityMembers(ctx context.Context, communityID CommunityID) ([]*Entity, error)
}
```

---

## 9. HTTP API

### 知识管理（文档层）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge | 创建知识 |
| GET | /api/knowledge | 列出知识 |
| GET | /api/knowledge/:id | 获取知识 |
| PUT | /api/knowledge/:id | 更新元信息 |
| POST | /api/knowledge/:id/archive | 归档知识 |

### 版本管理

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/versions | 创建版本 |
| GET | /api/knowledge/:id/versions | 列出版本 |
| GET | /api/knowledge/:id/versions/latest | 最新版本 |

### Chunk 存储（Agent 写入预切分数据）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/chunks | 批量写入 Chunk |
| GET | /api/knowledge/:id/chunks?version=N | 列出 Chunk |
| DELETE | /api/knowledge/:id/chunks | 删除 Chunk |

### Embedding 存储（Agent 写入预计算向量）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/embeddings | 批量写入 Embedding |

### 向量搜索（Agent 传入预计算的 query 向量）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/search | 向量相似度搜索 |

### 应用操作

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/publish | 发布知识 |
| POST | /api/knowledge/:id/rebuild-index | 清除旧索引数据 |

### Episode 管理（图谱层 - 数据溯源）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/episodes | 创建 Episode |
| GET | /api/episodes | 列出 Episode |
| GET | /api/episodes/:id | 获取 Episode |
| POST | /api/episodes/:id/mentions | 关联 Episode-Entity |
| GET | /api/episodes/:id/mentions | 列出 Episode 关联的实体 |

### Entity 管理（图谱层 - 实体）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/entities | 创建实体 |
| POST | /api/entities/batch | 批量创建实体 |
| GET | /api/entities | 列出实体 |
| GET | /api/entities/:id | 获取实体 |
| PUT | /api/entities/:id | 更新实体 |
| DELETE | /api/entities/:id | 删除实体 |
| POST | /api/entities/search | 语义搜索实体 |

### EntityEdge 管理（图谱层 - 实体关系）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/entity-edges | 创建实体关系 |
| POST | /api/entity-edges/batch | 批量创建实体关系 |
| GET | /api/entity-edges/:id | 获取实体关系 |
| PUT | /api/entity-edges/:id/invalidate | 标记实体关系失效 |
| GET | /api/entities/:id/edges | 列出实体的关系 |
| POST | /api/entity-edges/search | 语义搜索关系事实 |

### Community 管理（图谱层 - 社区聚类）

| 方法 | 路径 | 描述 |
|------|------|------|
| PUT | /api/communities/rebuild | 重建所有社区（全量替换） |
| GET | /api/communities | 列出社区 |
| GET | /api/communities/:id | 获取社区详情 |
| GET | /api/communities/:id/members | 列出社区成员 |

---

## 10. 职责边界

mygo（Data Plane）只负责：

- 数据的存储和查询（文档层 + 图谱层）
- 向量的持久化和相似度搜索（Chunk 向量 + Entity 向量 + Fact 向量）
- 状态管理和生命周期
- 双时态数据的存取和 point-in-time 查询

ave_mujica（Intelligence Plane）负责：

- Markdown 切分为 Chunk
- Chunk / Entity / Fact 生成 Embedding 向量
- **实体抽取**（从文档中发现 Entity）
- **关系发现**（从文档中发现 EntityEdge 三元组）
- **实体去重/合并**（判断两个实体是否为同一实体）
- **矛盾检测**（判断新边是否与旧边冲突，触发 invalidation）
- **社区检测**（运行 Leiden 等算法，将结果写入 mygo）
- Query 文本生成 Embedding 向量后调用 mygo 搜索 API
- Memory 管理（store/retrieve/reflect/forget）
- Agent 编排 / LLM 调用 / RAG pipeline
