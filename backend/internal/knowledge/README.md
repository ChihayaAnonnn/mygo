# Knowledge 领域模块

知识基础设施层：结构化存储、版本管理、Chunk 切分、Embedding 生成、图谱关系、语义检索。

作为 Agent External Cognitive System 的底层，为 Memory 模块提供存储和检索能力。

## 目录结构

```
knowledge/
├── domain/
│   ├── model.go        # Node, Version, Chunk, Edge, Embedding
│   ├── repository.go   # 5 个 Repository 接口
│   ├── service.go      # 6 个 Service 接口
│   └── types.go        # ID 类型、枚举、Command/Query、错误
│
├── application/
│   └── app_service.go  # KnowledgeApplicationService 实现
│
├── infrastructure/
│   ├── persistence/    # PO 模型（Node/Version/Chunk/Edge/Embedding）
│   ├── filesystem/     # 文件存储（预留）
│   └── ai/             # Embedder, Chunker（预留）
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
| `MarkdownRenderService` | 文件派生服务：DB → FS |
| `KnowledgeChunkService` | 切分服务：Markdown → Chunk |
| `EmbeddingService` | 向量计算：Chunk → Vector |
| `RetrievalService` | 语义检索：RAG 入口 |

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

### 应用操作

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/knowledge/:id/publish | 发布知识 |
| POST | /api/knowledge/:id/rebuild-index | 重建索引 |
| POST | /api/knowledge/search | 语义搜索 |

## 相关文档

- [Schema 设计](../../docs/knowledge_schema_design.md)
- [接口设计](../../docs/knowledge_interface_design.md)
