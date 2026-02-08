package persistence

import (
	"context"
	"errors"

	"mygo/internal/infra"
	"mygo/internal/knowledge/domain"
)

// ChunkRepository 知识分块仓储实现
type ChunkRepository struct {
	db *infra.GormDB
}

// NewChunkRepository 构造函数
func NewChunkRepository(res *infra.Resources) (*ChunkRepository, error) {
	if res == nil {
		return nil, errors.New("chunk repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("chunk repo: resources db is nil")
	}
	return &ChunkRepository{db: res.DB}, nil
}

// BatchCreate 批量创建分块
func (r *ChunkRepository) BatchCreate(ctx context.Context, chunks []*domain.Chunk) error {
	if len(chunks) == 0 {
		return nil
	}

	pos := make([]ChunkPO, 0, len(chunks))
	for _, c := range chunks {
		pos = append(pos, *ChunkPOFromDomain(c))
	}

	if err := r.db.WithContext(ctx).Create(&pos).Error; err != nil {
		return err
	}

	// 回写生成的 ID
	for i := range pos {
		*chunks[i] = *pos[i].ToDomain()
	}
	return nil
}

// ListByNodeVersion 按 node_id + version 列出分块
func (r *ChunkRepository) ListByNodeVersion(ctx context.Context, nodeID domain.KnowledgeID, version int) ([]*domain.Chunk, error) {
	if nodeID == "" {
		return nil, errors.New("chunk repo: node_id is required")
	}

	var pos []ChunkPO
	if err := r.db.WithContext(ctx).
		Where("node_id = ? AND version = ?", string(nodeID), version).
		Order("chunk_index ASC").
		Find(&pos).Error; err != nil {
		return nil, err
	}

	chunks := make([]*domain.Chunk, 0, len(pos))
	for i := range pos {
		chunks = append(chunks, pos[i].ToDomain())
	}
	return chunks, nil
}

// DeleteByNodeVersion 按 node_id + version 删除分块
func (r *ChunkRepository) DeleteByNodeVersion(ctx context.Context, nodeID domain.KnowledgeID, version int) error {
	if nodeID == "" {
		return errors.New("chunk repo: node_id is required")
	}

	return r.db.WithContext(ctx).
		Where("node_id = ? AND version = ?", string(nodeID), version).
		Delete(&ChunkPO{}).Error
}

// GetByIDs 根据 ChunkID 列表批量查找 Chunk
func (r *ChunkRepository) GetByIDs(ctx context.Context, chunkIDs []domain.ChunkID) ([]*domain.Chunk, error) {
	if len(chunkIDs) == 0 {
		return nil, nil
	}

	ids := make([]string, 0, len(chunkIDs))
	for _, id := range chunkIDs {
		ids = append(ids, string(id))
	}

	var pos []ChunkPO
	if err := r.db.WithContext(ctx).Where("chunk_id IN ?", ids).Find(&pos).Error; err != nil {
		return nil, err
	}

	// 按 chunkIDs 的原始顺序排列结果
	poMap := make(map[string]*ChunkPO, len(pos))
	for i := range pos {
		poMap[pos[i].ChunkID] = &pos[i]
	}

	chunks := make([]*domain.Chunk, 0, len(chunkIDs))
	for _, id := range chunkIDs {
		if po, ok := poMap[string(id)]; ok {
			chunks = append(chunks, po.ToDomain())
		}
	}

	return chunks, nil
}

// 编译时检查
var _ domain.ChunkRepository = (*ChunkRepository)(nil)
