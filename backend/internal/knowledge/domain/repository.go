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

// ==================== EdgeRepository ====================

// EdgeRepository 知识关系的仓储接口（基础设施层实现）
type EdgeRepository interface {
	Create(ctx context.Context, edge *Edge) error
	Delete(ctx context.Context, id string) error
	ListByFromNode(ctx context.Context, fromNodeID KnowledgeID) ([]*Edge, error)
	ListByToNode(ctx context.Context, toNodeID KnowledgeID) ([]*Edge, error)
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
