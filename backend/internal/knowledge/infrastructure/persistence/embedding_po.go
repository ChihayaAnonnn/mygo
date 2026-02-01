package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"

	"github.com/pgvector/pgvector-go"
)

// EmbeddingPO 是 knowledge_embeddings 的数据库存储模型（Persistence Object）
// 注意：使用 pgvector 扩展存储向量数据
type EmbeddingPO struct {
	ChunkID   string          `gorm:"column:chunk_id;type:uuid;primaryKey"`
	Embedding pgvector.Vector `gorm:"column:embedding;type:vector(1536);not null"`
	Model     string          `gorm:"column:model;type:varchar(64);not null"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
}

func (EmbeddingPO) TableName() string { return "knowledge_embeddings" }

// EmbeddingPOFromDomain 从领域模型转换为 PO
func EmbeddingPOFromDomain(e *domain.Embedding) *EmbeddingPO {
	if e == nil {
		return nil
	}
	return &EmbeddingPO{
		ChunkID:   e.ChunkID,
		Embedding: pgvector.NewVector(e.Embedding),
		Model:     e.Model,
		CreatedAt: e.CreatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *EmbeddingPO) ToDomain() *domain.Embedding {
	if p == nil {
		return nil
	}
	return &domain.Embedding{
		ChunkID:   p.ChunkID,
		Embedding: p.Embedding.Slice(),
		Model:     p.Model,
		CreatedAt: p.CreatedAt,
	}
}
