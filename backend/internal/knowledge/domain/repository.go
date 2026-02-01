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
	ListByVersion(ctx context.Context, versionID VersionID) ([]*Chunk, error)
	DeleteByVersion(ctx context.Context, versionID VersionID) error
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
