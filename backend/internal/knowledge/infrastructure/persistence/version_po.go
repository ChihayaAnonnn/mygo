package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"
)

// VersionPO 是 knowledge_versions 的数据库存储模型（Persistence Object）
type VersionPO struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey"`
	NodeID    string    `gorm:"column:node_id;type:uuid;not null;index:idx_knowledge_versions_node"`
	Version   int       `gorm:"column:version;not null"`
	ContentMd string    `gorm:"column:content_md;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime"`
}

func (VersionPO) TableName() string { return "knowledge_versions" }

// VersionPOFromDomain 从领域模型转换为 PO
func VersionPOFromDomain(e *domain.Version) *VersionPO {
	if e == nil {
		return nil
	}
	return &VersionPO{
		ID:        e.ID,
		NodeID:    e.NodeID,
		Version:   e.Version,
		ContentMd: e.ContentMd,
		CreatedAt: e.CreatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *VersionPO) ToDomain() *domain.Version {
	if p == nil {
		return nil
	}
	return &domain.Version{
		ID:        p.ID,
		NodeID:    p.NodeID,
		Version:   p.Version,
		ContentMd: p.ContentMd,
		CreatedAt: p.CreatedAt,
	}
}
