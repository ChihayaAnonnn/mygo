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

// ==================== 4. EpisodeService（数据摄入事件服务） ====================

// EpisodeService 数据摄入事件服务接口
// 职责：管理 Episode 的生命周期和 Episode-Entity 关联（数据溯源）
type EpisodeService interface {
	// CreateEpisode 创建 Episode
	CreateEpisode(ctx context.Context, cmd CreateEpisodeCmd) (EpisodeID, error)

	// GetEpisode 获取 Episode
	GetEpisode(ctx context.Context, id EpisodeID) (*Episode, error)

	// ListEpisodes 列出 Episode
	ListEpisodes(ctx context.Context, query EpisodeQuery) ([]*Episode, error)

	// AddMentions 关联 Episode 与 Entity（MENTIONS 关系）
	AddMentions(ctx context.Context, episodeID EpisodeID, entityIDs []EntityID) error

	// ListMentionedEntities 列出 Episode 关联的实体
	ListMentionedEntities(ctx context.Context, episodeID EpisodeID) ([]*Entity, error)
}

// ==================== 5. EntityService（实体服务） ====================

// EntityService 实体服务接口
// 职责：管理从文档/对话中抽取的细粒度实体，提供语义搜索能力
type EntityService interface {
	// CreateEntity 创建实体
	CreateEntity(ctx context.Context, cmd CreateEntityCmd) (EntityID, error)

	// BatchCreateEntities 批量创建实体
	BatchCreateEntities(ctx context.Context, cmds []CreateEntityCmd) ([]EntityID, error)

	// GetEntity 获取实体
	GetEntity(ctx context.Context, id EntityID) (*Entity, error)

	// UpdateEntity 更新实体
	UpdateEntity(ctx context.Context, cmd UpdateEntityCmd) error

	// DeleteEntity 删除实体
	DeleteEntity(ctx context.Context, id EntityID) error

	// ListEntities 列出实体
	ListEntities(ctx context.Context, query EntityQuery) ([]*Entity, error)

	// SearchEntities 通过语义向量搜索相似实体
	SearchEntities(ctx context.Context, embedding EmbeddingVector, topK int) ([]*Entity, error)
}

// ==================== 6. EntityEdgeService（实体关系服务） ====================

// EntityEdgeService 实体关系服务接口
// 职责：管理实体间的三元组关系，支持双时态查询和边失效
type EntityEdgeService interface {
	// CreateEntityEdge 创建实体关系
	CreateEntityEdge(ctx context.Context, cmd CreateEntityEdgeCmd) (EntityEdgeID, error)

	// BatchCreateEntityEdges 批量创建实体关系
	BatchCreateEntityEdges(ctx context.Context, cmds []CreateEntityEdgeCmd) ([]EntityEdgeID, error)

	// GetEntityEdge 获取实体关系
	GetEntityEdge(ctx context.Context, id EntityEdgeID) (*EntityEdge, error)

	// InvalidateEntityEdge 标记实体关系为失效（双时态：不删除，记录失效时间）
	InvalidateEntityEdge(ctx context.Context, id EntityEdgeID) error

	// ListEdgesByEntity 列出实体相关的所有关系
	ListEdgesByEntity(ctx context.Context, entityID EntityID) ([]*EntityEdge, error)

	// SearchEntityEdges 通过语义向量搜索相似的事实关系
	SearchEntityEdges(ctx context.Context, embedding EmbeddingVector, topK int) ([]*EntityEdge, error)
}

// ==================== 7. CommunityService（社区服务） ====================

// CommunityService 社区服务接口
// 职责：存储和查询社区检测结果（检测算法由 ave_mujica 执行）
type CommunityService interface {
	// RebuildCommunities 全量重建社区（删除旧社区，写入新社区及成员关系）
	RebuildCommunities(ctx context.Context, communities []CreateCommunityCmd) error

	// GetCommunity 获取社区
	GetCommunity(ctx context.Context, id CommunityID) (*Community, error)

	// ListCommunities 列出社区
	ListCommunities(ctx context.Context, query CommunityQuery) ([]*Community, error)

	// ListCommunityMembers 列出社区的成员实体
	ListCommunityMembers(ctx context.Context, communityID CommunityID) ([]*Entity, error)
}
