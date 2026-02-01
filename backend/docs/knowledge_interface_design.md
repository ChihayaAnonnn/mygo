# 核心 Service 接口设计（Go）

> 本文档定义核心 Service 接口（接口即边界），用于支撑 **Markdown 知识管理 → 文件派生 → AI 计算** 的完整闭环。
>
> 设计目标：
>
> - 接口即架构（Interface = Architecture）
> - 明确职责边界，避免 Service 膨胀
> - 为未来 AI / RAG / 多端消费预留稳定契约

---

## 1. KnowledgeService（知识元服务）

**职责**：

- 管理“知识”这一一等实体（Knowledge）
- 不关心具体内容，只关心元信息与生命周期

```go
type KnowledgeService interface {
    CreateKnowledge(ctx context.Context, cmd CreateKnowledgeCmd) (KnowledgeID, error)
    UpdateKnowledgeMeta(ctx context.Context, cmd UpdateKnowledgeMetaCmd) error
    GetKnowledge(ctx context.Context, id KnowledgeID) (Knowledge, error)
    ListKnowledge(ctx context.Context, query KnowledgeQuery) ([]Knowledge, error)
    ArchiveKnowledge(ctx context.Context, id KnowledgeID) error
}
```

**说明**：

- `Knowledge` 是**逻辑知识单元**，而不是某一篇 Markdown
- 删除通常是软删除（Archive），以保证历史可追溯

---

## 2. KnowledgeVersionService（版本服务）

**职责**：

- 管理 Markdown 内容的版本演化
- 数据库中的 `knowledge_versions` 是事实源

```go
type KnowledgeVersionService interface {
    CreateVersion(ctx context.Context, cmd CreateVersionCmd) (VersionID, error)
    GetVersion(ctx context.Context, id VersionID) (KnowledgeVersion, error)
    GetLatestVersion(ctx context.Context, knowledgeID KnowledgeID) (KnowledgeVersion, error)
    ListVersions(ctx context.Context, knowledgeID KnowledgeID) ([]KnowledgeVersion, error)
}
```

**说明**：

- 每一次内容修改 = 一个新 Version（不可变）
- `content_md` 永远以 DB 为准

---

## 3. MarkdownRenderService（文件派生服务）

**职责**：

- 将 DB 中的 Markdown 内容**派生**为文件系统产物
- FS 永远不是事实源

```go
type MarkdownRenderService interface {
    RenderToFile(ctx context.Context, version KnowledgeVersion) (FilePath, error)
    RemoveFile(ctx context.Context, versionID VersionID) error
}
```

**说明**：

- 典型使用场景：
  - 本地预览
  - Git 同步
  - 静态站点生成

---

## 4. KnowledgeChunkService（切分服务）

**职责**：

- 将 Markdown 内容切分为语义 Chunk
- 为 AI / Embedding 提供最小语义单元

```go
type KnowledgeChunkService interface {
    BuildChunks(ctx context.Context, version KnowledgeVersion) ([]KnowledgeChunk, error)
    ListChunks(ctx context.Context, versionID VersionID) ([]KnowledgeChunk, error)
}
```

**说明**：

- Chunk 是**可重建派生数据**
- 允许未来切分策略变更（不影响原始内容）

---

## 5. EmbeddingService（向量计算服务）

**职责**：

- 调用 Python / 模型服务生成向量
- 只关心输入输出，不关心存储

```go
type EmbeddingService interface {
    EmbedChunks(ctx context.Context, chunks []KnowledgeChunk) ([]EmbeddingVector, error)
}
```

**说明**：

- Go 负责调度
- Python 负责模型计算
- 可替换模型、可多后端

---

## 6. RetrievalService（检索服务）

**职责**：

- 基于向量进行语义检索
- 是 RAG / 问答系统的入口

```go
type RetrievalService interface {
    Search(ctx context.Context, query string, topK int) ([]KnowledgeChunk, error)
}
```

**说明**：

- 本质是“知识消费接口”
- 不直接暴露 Version / Knowledge 细节

---

## 7. KnowledgeApplicationService（应用层编排）

**职责**：

- 编排多个底层 Service
- 对外提供高层用例（Use Case）

```go
type KnowledgeApplicationService interface {
    PublishKnowledge(ctx context.Context, knowledgeID KnowledgeID) error
    RebuildIndex(ctx context.Context, knowledgeID KnowledgeID) error
}
```

**说明**：

- 这里是**业务流程**而非领域逻辑
- 非常适合做 Saga / Workflow

---
