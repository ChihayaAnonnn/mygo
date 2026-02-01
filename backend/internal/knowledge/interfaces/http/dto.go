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

// ==================== Search DTOs ====================

// SearchRequest 搜索请求
type SearchRequest struct {
	Query string `json:"query" binding:"required"`
	TopK  int    `json:"top_k"`
}
