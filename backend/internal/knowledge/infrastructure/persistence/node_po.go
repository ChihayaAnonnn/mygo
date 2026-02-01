package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"
)

// NodePO 是 knowledge_nodes 的数据库存储模型（Persistence Object）
type NodePO struct {
	ID             string    `gorm:"column:id;type:uuid;primaryKey"`
	NodeType       string    `gorm:"column:node_type;type:varchar(32);not null;index:idx_knowledge_nodes_type"`
	Title          string    `gorm:"column:title;type:text;not null"`
	Summary        string    `gorm:"column:summary;type:text"`
	Status         string    `gorm:"column:status;type:varchar(16);default:'draft';index:idx_knowledge_nodes_status"`
	Confidence     *float32  `gorm:"column:confidence;type:real"`
	CurrentVersion int       `gorm:"column:current_version;not null;default:1"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (NodePO) TableName() string { return "knowledge_nodes" }

// NodePOFromDomain 从领域模型转换为 PO
func NodePOFromDomain(e *domain.Node) *NodePO {
	if e == nil {
		return nil
	}
	return &NodePO{
		ID:             e.ID,
		NodeType:       string(e.NodeType),
		Title:          e.Title,
		Summary:        e.Summary,
		Status:         string(e.Status),
		Confidence:     e.Confidence,
		CurrentVersion: e.CurrentVersion,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *NodePO) ToDomain() *domain.Node {
	if p == nil {
		return nil
	}
	return &domain.Node{
		ID:             p.ID,
		NodeType:       domain.NodeType(p.NodeType),
		Title:          p.Title,
		Summary:        p.Summary,
		Status:         domain.NodeStatus(p.Status),
		Confidence:     p.Confidence,
		CurrentVersion: p.CurrentVersion,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}
