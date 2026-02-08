# Knowledge 领域模块

Data Plane 数据服务：为外部 Agent 提供 Knowledge 数据的 CRUD 和向量检索 API。

## 目录结构

```text
knowledge/
├── domain/
│   ├── model.go        # Node, Version, Chunk, Edge, Embedding
│   ├── repository.go   # Repository 接口
│   ├── service.go      # Service 接口（KnowledgeService, VersionService）
│   ├── types.go        # ID 类型、枚举、Command/Query、错误
│   ├── ai_task.go      # AITask 模型
│   └── tag.go          # Tag 模型
│
├── application/
│   ├── app_service.go      # KnowledgeApplicationService（发布、重建索引）
│   ├── knowledge_service.go # KnowledgeService 实现
│   ├── version_service.go   # VersionService 实现
│   └── ai_task_service.go   # AITaskService 实现
│
├── infrastructure/
│   └── persistence/    # PO 模型 + Repository 实现
│
└── interfaces/http/
    ├── handler.go
    ├── routes.go
    └── dto.go
```

## 领域服务

| 接口 | 职责 |
|------|------|
| `KnowledgeService` | 知识元服务：元信息与生命周期 |
| `KnowledgeVersionService` | 版本服务：Markdown 版本演化 |
| `MarkdownRenderService` | 文件派生服务：DB -> FS（预留） |

## API 接口

### 知识管理

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

### Chunk 管理（Agent 写入预切分数据）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/chunks | 批量写入 Chunk |
| GET | /api/knowledge/:id/chunks | 列出 Chunk |
| DELETE | /api/knowledge/:id/chunks | 删除 Chunk |

### Embedding 管理（Agent 写入预计算向量）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/embeddings | 批量写入 Embedding |

### 搜索（Agent 传入预计算的 query 向量）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/search | 向量相似度搜索 |

### 应用操作

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/publish | 发布知识（更新状态） |
| POST | /api/knowledge/:id/rebuild-index | 清除旧索引数据 |

## 相关文档

- [Schema 设计](../../docs/knowledge_schema_design.md)
- [接口设计](../../docs/knowledge_interface_design.md)
