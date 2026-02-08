package persistence

import (
	"context"
	"errors"

	"mygo/internal/infra"
	"mygo/internal/knowledge/domain"

	"gorm.io/gorm"
)

// VersionRepository 知识版本仓储实现
type VersionRepository struct {
	db *infra.GormDB
}

// NewVersionRepository 构造函数
func NewVersionRepository(res *infra.Resources) (*VersionRepository, error) {
	if res == nil {
		return nil, errors.New("version repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("version repo: resources db is nil")
	}
	return &VersionRepository{db: res.DB}, nil
}

// Create 创建知识版本
func (r *VersionRepository) Create(ctx context.Context, version *domain.Version) error {
	if version == nil {
		return errors.New("version repo: version is nil")
	}
	if version.VersionID == "" {
		return errors.New("version repo: version_id is required")
	}
	if version.NodeID == "" {
		return errors.New("version repo: node_id is required")
	}

	p := VersionPOFromDomain(version)
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		return err
	}

	*version = *p.ToDomain()
	return nil
}

// GetByID 根据 version_id 获取版本
func (r *VersionRepository) GetByID(ctx context.Context, id domain.VersionID) (*domain.Version, error) {
	if id == "" {
		return nil, errors.New("version repo: version_id is required")
	}

	var po VersionPO
	if err := r.db.WithContext(ctx).Where("version_id = ?", string(id)).First(&po).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrVersionNotFound
		}
		return nil, err
	}

	return po.ToDomain(), nil
}

// GetLatestByKnowledge 获取知识节点的最新版本
func (r *VersionRepository) GetLatestByKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) (*domain.Version, error) {
	if knowledgeID == "" {
		return nil, errors.New("version repo: knowledge_id is required")
	}

	var po VersionPO
	if err := r.db.WithContext(ctx).
		Where("node_id = ?", string(knowledgeID)).
		Order("version DESC").
		First(&po).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrVersionNotFound
		}
		return nil, err
	}

	return po.ToDomain(), nil
}

// ListByKnowledge 列出知识节点的所有版本
func (r *VersionRepository) ListByKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) ([]*domain.Version, error) {
	if knowledgeID == "" {
		return nil, errors.New("version repo: knowledge_id is required")
	}

	var pos []VersionPO
	if err := r.db.WithContext(ctx).
		Where("node_id = ?", string(knowledgeID)).
		Order("version DESC").
		Find(&pos).Error; err != nil {
		return nil, err
	}

	versions := make([]*domain.Version, 0, len(pos))
	for i := range pos {
		versions = append(versions, pos[i].ToDomain())
	}
	return versions, nil
}

// 编译时检查
var _ domain.VersionRepository = (*VersionRepository)(nil)
