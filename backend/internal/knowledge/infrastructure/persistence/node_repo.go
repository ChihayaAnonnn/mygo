package persistence

import (
	"context"
	"errors"

	"mygo/internal/infra"
	"mygo/internal/knowledge/domain"

	"gorm.io/gorm"
)

// NodeRepository 知识节点仓储实现
type NodeRepository struct {
	db *infra.GormDB
}

// NewNodeRepository 构造函数
func NewNodeRepository(res *infra.Resources) (*NodeRepository, error) {
	if res == nil {
		return nil, errors.New("node repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("node repo: resources db is nil")
	}
	return &NodeRepository{db: res.DB}, nil
}

// Create 创建知识节点
func (r *NodeRepository) Create(ctx context.Context, node *domain.Node) error {
	if r.db == nil {
		return errors.New("node repo: db is nil")
	}
	if node == nil {
		return errors.New("node repo: node is nil")
	}
	// 必填字段校验
	if node.NodeID == "" {
		return errors.New("node repo: node_id is required")
	}
	if node.NodeType == "" {
		return errors.New("node repo: node_type is required")
	}
	if node.Title == "" {
		return errors.New("node repo: title is required")
	}
	// Summary、Confidence 是可选字段，不校验
	// CreatedAt、UpdatedAt 由数据库自动生成，不校验

	p := NodePOFromDomain(node)
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		return err
	}

	// GORM 会自动将数据库生成的自增ID填充到 p.ID
	// 将完整的PO（包含ID、CreatedAt、UpdatedAt）转回领域模型
	*node = *p.ToDomain()
	return nil
}

// GetByID 根据ID获取知识节点
func (r *NodeRepository) GetByID(ctx context.Context, id domain.KnowledgeID) (*domain.Node, error) {
	if id == "" {
		return nil, errors.New("node repo: node_id is required")
	}

	var po NodePO
	if err := r.db.WithContext(ctx).Where("node_id = ?", string(id)).First(&po).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrKnowledgeNotFound
		}
		return nil, err
	}

	return po.ToDomain(), nil
}

// Update 更新知识节点
func (r *NodeRepository) Update(ctx context.Context, node *domain.Node) error {
	if node == nil {
		return errors.New("node repo: node is nil")
	}
	if node.NodeID == "" {
		return errors.New("node repo: node_id is required")
	}

	p := NodePOFromDomain(node)
	result := r.db.WithContext(ctx).Model(&NodePO{}).Where("node_id = ?", p.NodeID).Updates(map[string]any{
		"node_type":       p.NodeType,
		"title":           p.Title,
		"summary":         p.Summary,
		"status":          p.Status,
		"confidence":      p.Confidence,
		"current_version": p.CurrentVersion,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrKnowledgeNotFound
	}

	return nil
}

// Delete 删除知识节点
func (r *NodeRepository) Delete(ctx context.Context, id domain.KnowledgeID) error {
	if id == "" {
		return errors.New("node repo: node_id is required")
	}

	result := r.db.WithContext(ctx).Where("node_id = ?", string(id)).Delete(&NodePO{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrKnowledgeNotFound
	}

	return nil
}

// List 查询知识节点列表
func (r *NodeRepository) List(ctx context.Context, query domain.KnowledgeQuery) ([]*domain.Node, error) {
	db := r.db.WithContext(ctx).Model(&NodePO{})

	if query.NodeType != nil {
		db = db.Where("node_type = ?", string(*query.NodeType))
	}
	if query.Status != nil {
		db = db.Where("status = ?", string(*query.Status))
	}
	if query.Keyword != nil && *query.Keyword != "" {
		like := "%" + *query.Keyword + "%"
		db = db.Where("title ILIKE ? OR summary ILIKE ?", like, like)
	}

	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	} else {
		db = db.Limit(20)
	}
	if query.Offset > 0 {
		db = db.Offset(query.Offset)
	}

	db = db.Order("created_at DESC")

	var pos []NodePO
	if err := db.Find(&pos).Error; err != nil {
		return nil, err
	}

	nodes := make([]*domain.Node, 0, len(pos))
	for i := range pos {
		nodes = append(nodes, pos[i].ToDomain())
	}
	return nodes, nil
}

// 编译时检查：确保 NodeRepository 实现了 domain.KnowledgeRepository 接口
// 这是 Go 的最佳实践，如果接口方法有变化，编译时会报错
var _ domain.KnowledgeRepository = (*NodeRepository)(nil)
