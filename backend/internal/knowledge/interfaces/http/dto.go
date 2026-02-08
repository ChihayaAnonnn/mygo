package http

import "time"

// ==================== Knowledge DTOs ====================

// CreateKnowledgeRequest 创建知识请求
type CreateKnowledgeRequest struct {
	NodeType string `json:"node_type" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Summary  string `json:"summary"`
}

// UpdateKnowledgeMetaRequest 更新知识元信息请求
type UpdateKnowledgeMetaRequest struct {
	Title      *string  `json:"title"`
	Summary    *string  `json:"summary"`
	Status     *string  `json:"status"`
	Confidence *float32 `json:"confidence"`
}

// KnowledgeResponse 知识响应
type KnowledgeResponse struct {
	ID             string    `json:"id"`
	NodeType       string    `json:"node_type"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	Status         string    `json:"status"`
	Confidence     *float32  `json:"confidence,omitempty"`
	CurrentVersion int       `json:"current_version"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ListKnowledgeRequest 列表请求
type ListKnowledgeRequest struct {
	NodeType string `form:"node_type"`
	Status   string `form:"status"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

// ==================== Version DTOs ====================

// CreateVersionRequest 创建版本请求
type CreateVersionRequest struct {
	ContentMd string `json:"content_md" binding:"required"`
}

// VersionResponse 版本响应
type VersionResponse struct {
	ID        string    `json:"id"`
	NodeID    string    `json:"node_id"`
	Version   int       `json:"version"`
	ContentMd string    `json:"content_md"`
	CreatedAt time.Time `json:"created_at"`
}

// ==================== Chunk DTOs ====================

// BatchCreateChunksRequest 批量创建分块请求（Agent 端预切分后写入）
type BatchCreateChunksRequest struct {
	Version int               `json:"version" binding:"required"`
	Chunks  []CreateChunkItem `json:"chunks" binding:"required,min=1"`
}

// CreateChunkItem 单个分块数据
type CreateChunkItem struct {
	ChunkID     string `json:"chunk_id" binding:"required"`
	HeadingPath string `json:"heading_path"`
	Content     string `json:"content" binding:"required"`
	TokenCount  *int   `json:"token_count"`
	ChunkIndex  *int   `json:"chunk_index"`
}

// DeleteChunksRequest 删除分块请求
type DeleteChunksRequest struct {
	Version int `json:"version" binding:"required"`
}

// ChunkResponse 分块响应
type ChunkResponse struct {
	ID          string    `json:"id"`
	NodeID      string    `json:"node_id"`
	Version     int       `json:"version"`
	HeadingPath string    `json:"heading_path"`
	Content     string    `json:"content"`
	TokenCount  *int      `json:"token_count,omitempty"`
	ChunkIndex  *int      `json:"chunk_index,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ==================== Embedding DTOs ====================

// BatchCreateEmbeddingsRequest 批量创建向量请求（Agent 端预计算后写入）
type BatchCreateEmbeddingsRequest struct {
	Embeddings []CreateEmbeddingItem `json:"embeddings" binding:"required,min=1"`
}

// CreateEmbeddingItem 单个向量数据
type CreateEmbeddingItem struct {
	ChunkID   string    `json:"chunk_id" binding:"required"`
	Embedding []float32 `json:"embedding" binding:"required"`
	Model     string    `json:"model" binding:"required"`
}

// ==================== Search DTOs ====================

// SearchRequest 向量搜索请求（Agent 端传入预计算的 query 向量）
type SearchRequest struct {
	Vector []float32 `json:"vector" binding:"required"`
	TopK   int       `json:"top_k"`
}
