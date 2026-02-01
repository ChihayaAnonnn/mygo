package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"
)

// EdgePO 是 knowledge_edges 的数据库存储模型（Persistence Object）
type EdgePO struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey"`
	FromNode  string    `gorm:"column:from_node;type:uuid;not null;index:idx_knowledge_edges_from"`
	ToNode    string    `gorm:"column:to_node;type:uuid;not null;index:idx_knowledge_edges_to"`
	EdgeType  string    `gorm:"column:edge_type;type:varchar(64);not null;index:idx_knowledge_edges_type"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (EdgePO) TableName() string { return "knowledge_edges" }

// EdgePOFromDomain 从领域模型转换为 PO
func EdgePOFromDomain(e *domain.Edge) *EdgePO {
	if e == nil {
		return nil
	}
	return &EdgePO{
		ID:        e.ID,
		FromNode:  e.FromNode,
		ToNode:    e.ToNode,
		EdgeType:  string(e.EdgeType),
		CreatedAt: e.CreatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *EdgePO) ToDomain() *domain.Edge {
	if p == nil {
		return nil
	}
	return &domain.Edge{
		ID:        p.ID,
		FromNode:  p.FromNode,
		ToNode:    p.ToNode,
		EdgeType:  domain.EdgeType(p.EdgeType),
		CreatedAt: p.CreatedAt,
	}
}
