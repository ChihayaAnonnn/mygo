package domain

import "context"

// ==================== KnowledgeRepository ====================

// KnowledgeRepository 知识节点的仓储接口（基础设施层实现）
type KnowledgeRepository interface {
	Create(ctx context.Context, node *Node) error
	GetByID(ctx context.Context, id KnowledgeID) (*Node, error)
	Update(ctx context.Context, node *Node) error
	Delete(ctx context.Context, id KnowledgeID) error
	List(ctx context.Context, query KnowledgeQuery) ([]*Node, error)
}

// ==================== VersionRepository ====================

// VersionRepository 知识版本的仓储接口（基础设施层实现）
type VersionRepository interface {
	Create(ctx context.Context, version *Version) error
	GetByID(ctx context.Context, id VersionID) (*Version, error)
	GetLatestByKnowledge(ctx context.Context, knowledgeID KnowledgeID) (*Version, error)
	ListByKnowledge(ctx context.Context, knowledgeID KnowledgeID) ([]*Version, error)
}

// ==================== ChunkRepository ====================

// ChunkRepository 知识分块的仓储接口（基础设施层实现）
type ChunkRepository interface {
	BatchCreate(ctx context.Context, chunks []*Chunk) error
	ListByNodeVersion(ctx context.Context, nodeID KnowledgeID, version int) ([]*Chunk, error)
	DeleteByNodeVersion(ctx context.Context, nodeID KnowledgeID, version int) error
}

// ==================== EmbeddingRepository ====================

// EmbeddingRepository 知识向量的仓储接口（基础设施层实现）
type EmbeddingRepository interface {
	BatchCreate(ctx context.Context, embeddings []*Embedding) error
	DeleteByChunkIDs(ctx context.Context, chunkIDs []ChunkID) error
	// SearchSimilar 向量相似度搜索，返回最相似的 topK 个 ChunkID
	SearchSimilar(ctx context.Context, embedding EmbeddingVector, topK int) ([]ChunkID, error)
}

// ==================== EdgeRepository（增强版）====================

// EdgeRepository 文档级关系的仓储接口（基础设施层实现）
type EdgeRepository interface {
	Create(ctx context.Context, edge *Edge) error
	GetByID(ctx context.Context, id string) (*Edge, error)
	Delete(ctx context.Context, id string) error
	// Invalidate 标记边为失效（双时态：记录 invalidated_at 而非删除）
	Invalidate(ctx context.Context, id string) error
	ListByFromNode(ctx context.Context, fromNodeID KnowledgeID) ([]*Edge, error)
	ListByToNode(ctx context.Context, toNodeID KnowledgeID) ([]*Edge, error)
	// ListActiveByNode 列出节点相关的所有有效边（未失效且在有效期内）
	ListActiveByNode(ctx context.Context, nodeID KnowledgeID) ([]*Edge, error)
}

// ==================== EpisodeRepository ====================

// EpisodeRepository 数据摄入事件的仓储接口（基础设施层实现）
type EpisodeRepository interface {
	Create(ctx context.Context, episode *Episode) error
	GetByID(ctx context.Context, id EpisodeID) (*Episode, error)
	List(ctx context.Context, query EpisodeQuery) ([]*Episode, error)
	// AddMentions 记录 Episode 中提到的实体
	AddMentions(ctx context.Context, episodeID EpisodeID, entityIDs []EntityID) error
	// ListMentionedEntityIDs 列出 Episode 关联的实体 ID
	ListMentionedEntityIDs(ctx context.Context, episodeID EpisodeID) ([]EntityID, error)
	// ListEpisodesByEntity 列出提及某实体的所有 Episode ID
	ListEpisodesByEntity(ctx context.Context, entityID EntityID) ([]EpisodeID, error)
}

// ==================== EntityRepository ====================

// EntityRepository 实体的仓储接口（基础设施层实现）
type EntityRepository interface {
	Create(ctx context.Context, entity *Entity) error
	BatchCreate(ctx context.Context, entities []*Entity) error
	GetByID(ctx context.Context, id EntityID) (*Entity, error)
	Update(ctx context.Context, entity *Entity) error
	Delete(ctx context.Context, id EntityID) error
	List(ctx context.Context, query EntityQuery) ([]*Entity, error)
	// SearchByEmbedding 通过语义向量搜索相似实体
	SearchByEmbedding(ctx context.Context, embedding EmbeddingVector, topK int) ([]*Entity, error)
}

// ==================== EntityEdgeRepository ====================

// EntityEdgeRepository 实体关系的仓储接口（基础设施层实现）
type EntityEdgeRepository interface {
	Create(ctx context.Context, edge *EntityEdge) error
	BatchCreate(ctx context.Context, edges []*EntityEdge) error
	GetByID(ctx context.Context, id EntityEdgeID) (*EntityEdge, error)
	// Invalidate 标记实体关系为失效（双时态）
	Invalidate(ctx context.Context, id EntityEdgeID) error
	// ListByEntity 列出实体相关的所有边（作为 Subject 或 Object）
	ListByEntity(ctx context.Context, entityID EntityID) ([]*EntityEdge, error)
	// ListActiveByEntity 列出实体相关的所有有效边
	ListActiveByEntity(ctx context.Context, entityID EntityID) ([]*EntityEdge, error)
	// SearchByEmbedding 通过语义向量搜索相似的事实关系
	SearchByEmbedding(ctx context.Context, embedding EmbeddingVector, topK int) ([]*EntityEdge, error)
}

// ==================== CommunityRepository ====================

// CommunityRepository 社区的仓储接口（基础设施层实现）
type CommunityRepository interface {
	Create(ctx context.Context, community *Community) error
	GetByID(ctx context.Context, id CommunityID) (*Community, error)
	List(ctx context.Context, query CommunityQuery) ([]*Community, error)
	// DeleteAll 删除所有社区及其成员关系（用于全量重建）
	DeleteAll(ctx context.Context) error
	// AddMembers 为社区添加成员实体
	AddMembers(ctx context.Context, communityID CommunityID, entityIDs []EntityID) error
	// ListMemberIDs 列出社区的成员实体 ID
	ListMemberIDs(ctx context.Context, communityID CommunityID) ([]EntityID, error)
	// ListMemberEntities 列出社区的成员实体（完整对象）
	ListMemberEntities(ctx context.Context, communityID CommunityID) ([]*Entity, error)
}

// ==================== AITaskRepository ====================

// AITaskRepository AI 任务的仓储接口（基础设施层实现）
type AITaskRepository interface {
	Create(ctx context.Context, task *AITask) error
	GetByID(ctx context.Context, id AITaskID) (*AITask, error)
	Update(ctx context.Context, task *AITask) error
	ListByNode(ctx context.Context, nodeID KnowledgeID) ([]*AITask, error)
	ListByStatus(ctx context.Context, status AITaskStatus) ([]*AITask, error)
	ListPending(ctx context.Context, limit int) ([]*AITask, error)
}

// ==================== TagRepository ====================

// TagRepository 标签的仓储接口（基础设施层实现）
type TagRepository interface {
	Create(ctx context.Context, tag *Tag) error
	GetByID(ctx context.Context, id TagID) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
	Delete(ctx context.Context, id TagID) error
	List(ctx context.Context) ([]*Tag, error)
	// GetOrCreate 获取或创建标签（幂等操作）
	GetOrCreate(ctx context.Context, name string) (*Tag, error)
}

// ==================== KnowledgeNodeTagRepository ====================

// KnowledgeNodeTagRepository 节点标签关联的仓储接口（基础设施层实现）
type KnowledgeNodeTagRepository interface {
	Create(ctx context.Context, nodeTag *KnowledgeNodeTag) error
	Delete(ctx context.Context, nodeID KnowledgeID, tagID TagID) error
	ListTagsByNode(ctx context.Context, nodeID KnowledgeID) ([]*Tag, error)
	ListNodesByTag(ctx context.Context, tagID TagID) ([]KnowledgeID, error)
	SetNodeTags(ctx context.Context, nodeID KnowledgeID, tagIDs []TagID) error
}
