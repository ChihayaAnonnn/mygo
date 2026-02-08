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
