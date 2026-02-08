package persistence

import (
	"context"
	"errors"
	"fmt"

	"mygo/internal/infra"
	"mygo/internal/knowledge/domain"

	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

// EmbeddingRepository 知识向量仓储实现
type EmbeddingRepository struct {
	db *infra.GormDB
}

// NewEmbeddingRepository 构造函数
func NewEmbeddingRepository(res *infra.Resources) (*EmbeddingRepository, error) {
	if res == nil {
		return nil, errors.New("embedding repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("embedding repo: resources db is nil")
	}
	return &EmbeddingRepository{db: res.DB}, nil
}

// BatchCreate 批量创建向量
func (r *EmbeddingRepository) BatchCreate(ctx context.Context, embeddings []*domain.Embedding) error {
	if len(embeddings) == 0 {
		return nil
	}

	pos := make([]EmbeddingPO, 0, len(embeddings))
	for _, e := range embeddings {
		pos = append(pos, *EmbeddingPOFromDomain(e))
	}

	if err := r.db.WithContext(ctx).Create(&pos).Error; err != nil {
		return err
	}

	for i := range pos {
		*embeddings[i] = *pos[i].ToDomain()
	}
	return nil
}

// DeleteByChunkIDs 按 ChunkID 列表删除向量
func (r *EmbeddingRepository) DeleteByChunkIDs(ctx context.Context, chunkIDs []domain.ChunkID) error {
	if len(chunkIDs) == 0 {
		return nil
	}

	ids := make([]string, 0, len(chunkIDs))
	for _, id := range chunkIDs {
		ids = append(ids, string(id))
	}

	return r.db.WithContext(ctx).
		Where("chunk_id IN ?", ids).
		Delete(&EmbeddingPO{}).Error
}

// SearchSimilar 向量相似度搜索（cosine distance），返回最相似的 topK 个 ChunkID
func (r *EmbeddingRepository) SearchSimilar(ctx context.Context, embedding domain.EmbeddingVector, topK int) ([]domain.ChunkID, error) {
	if len(embedding) == 0 {
		return nil, errors.New("embedding repo: embedding vector is empty")
	}
	if topK <= 0 {
		topK = 10
	}

	vec := pgvector.NewVector(embedding)

	var results []struct {
		ChunkID string `gorm:"column:chunk_id"`
	}

	// 使用 pgvector 的 cosine distance 操作符 <=>
	err := r.db.WithContext(ctx).
		Model(&EmbeddingPO{}).
		Select("chunk_id").
		Order(gorm.Expr(fmt.Sprintf("embedding <=> '%s'", vec))).
		Limit(topK).
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	chunkIDs := make([]domain.ChunkID, 0, len(results))
	for _, r := range results {
		chunkIDs = append(chunkIDs, domain.ChunkID(r.ChunkID))
	}
	return chunkIDs, nil
}

// 编译时检查
var _ domain.EmbeddingRepository = (*EmbeddingRepository)(nil)
