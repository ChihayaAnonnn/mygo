package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"
)

// ChunkPO 是 knowledge_chunks 的数据库存储模型（Persistence Object）
type ChunkPO struct {
	ID          string    `gorm:"column:id;type:uuid;primaryKey"`
	NodeID      string    `gorm:"column:node_id;type:uuid;not null;index:idx_knowledge_chunks_node"`
	Version     int       `gorm:"column:version;not null;index:idx_knowledge_chunks_version,priority:2"`
	HeadingPath string    `gorm:"column:heading_path;type:text"`
	Content     string    `gorm:"column:content;type:text;not null"`
	TokenCount  *int      `gorm:"column:token_count;type:int"`
	ChunkIndex  *int      `gorm:"column:chunk_index;type:int"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ChunkPO) TableName() string { return "knowledge_chunks" }

// ChunkPOFromDomain 从领域模型转换为 PO
func ChunkPOFromDomain(e *domain.Chunk) *ChunkPO {
	if e == nil {
		return nil
	}
	return &ChunkPO{
		ID:          e.ID,
		NodeID:      e.NodeID,
		Version:     e.Version,
		HeadingPath: e.HeadingPath,
		Content:     e.Content,
		TokenCount:  e.TokenCount,
		ChunkIndex:  e.ChunkIndex,
		CreatedAt:   e.CreatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *ChunkPO) ToDomain() *domain.Chunk {
	if p == nil {
		return nil
	}
	return &domain.Chunk{
		ID:          p.ID,
		NodeID:      p.NodeID,
		Version:     p.Version,
		HeadingPath: p.HeadingPath,
		Content:     p.Content,
		TokenCount:  p.TokenCount,
		ChunkIndex:  p.ChunkIndex,
		CreatedAt:   p.CreatedAt,
	}
}
