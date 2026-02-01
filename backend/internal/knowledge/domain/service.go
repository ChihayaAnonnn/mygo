package domain

import "context"

// ==================== 1. KnowledgeService（知识元服务） ====================

// KnowledgeService 知识元服务接口
// 职责：管理"知识"这一一等实体（Knowledge），不关心具体内容，只关心元信息与生命周期
type KnowledgeService interface {
	// CreateKnowledge 创建知识
	CreateKnowledge(ctx context.Context, cmd CreateKnowledgeCmd) (KnowledgeID, error)

	// UpdateKnowledgeMeta 更新知识元信息
	UpdateKnowledgeMeta(ctx context.Context, cmd UpdateKnowledgeMetaCmd) error

	// GetKnowledge 获取知识
	GetKnowledge(ctx context.Context, id KnowledgeID) (*Node, error)

	// ListKnowledge 列出知识
	ListKnowledge(ctx context.Context, query KnowledgeQuery) ([]*Node, error)

	// ArchiveKnowledge 归档知识（软删除）
	ArchiveKnowledge(ctx context.Context, id KnowledgeID) error
}

// ==================== 2. KnowledgeVersionService（版本服务） ====================

// KnowledgeVersionService 版本服务接口
// 职责：管理 Markdown 内容的版本演化，数据库中的 knowledge_versions 是事实源
type KnowledgeVersionService interface {
	// CreateVersion 创建新版本（每一次内容修改 = 一个新 Version，不可变）
	CreateVersion(ctx context.Context, cmd CreateVersionCmd) (VersionID, error)

	// GetVersion 获取指定版本
	GetVersion(ctx context.Context, id VersionID) (*Version, error)

	// GetLatestVersion 获取知识的最新版本
	GetLatestVersion(ctx context.Context, knowledgeID KnowledgeID) (*Version, error)

	// ListVersions 列出知识的所有版本
	ListVersions(ctx context.Context, knowledgeID KnowledgeID) ([]*Version, error)
}

// ==================== 3. MarkdownRenderService（文件派生服务） ====================

// MarkdownRenderService 文件派生服务接口
// 职责：将 DB 中的 Markdown 内容派生为文件系统产物（FS 永远不是事实源）
// 典型使用场景：本地预览、Git 同步、静态站点生成
type MarkdownRenderService interface {
	// RenderToFile 将版本内容渲染为文件
	RenderToFile(ctx context.Context, version *Version) (FilePath, error)

	// RemoveFile 移除版本对应的文件
	RemoveFile(ctx context.Context, versionID VersionID) error
}

// ==================== 4. KnowledgeChunkService（切分服务） ====================

// KnowledgeChunkService 切分服务接口
// 职责：将 Markdown 内容切分为语义 Chunk，为 AI / Embedding 提供最小语义单元
// Chunk 是可重建派生数据，允许未来切分策略变更（不影响原始内容）
type KnowledgeChunkService interface {
	// BuildChunks 构建版本的所有 Chunk
	BuildChunks(ctx context.Context, version *Version) ([]*Chunk, error)

	// ListChunks 列出版本的所有 Chunk
	ListChunks(ctx context.Context, versionID VersionID) ([]*Chunk, error)
}

// ==================== 5. EmbeddingService（向量计算服务） ====================

// EmbeddingService 向量计算服务接口
// 职责：调用 Python / 模型服务生成向量
type EmbeddingService interface {
	// EmbedChunks 将 Chunk 列表转换为向量
	EmbedChunks(ctx context.Context, chunks []*Chunk) ([]EmbeddingVector, error)
}

// ==================== 6. RetrievalService（检索服务） ====================

// RetrievalService 检索服务接口
// 职责：基于向量进行语义检索，是 RAG / 问答系统的入口
type RetrievalService interface {
	// Search 语义搜索，返回最相关的 Chunk
	Search(ctx context.Context, query string, topK int) ([]*Chunk, error)
}
