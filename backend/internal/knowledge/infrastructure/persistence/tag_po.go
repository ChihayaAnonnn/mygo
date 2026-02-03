package persistence

import (
	"mygo/internal/knowledge/domain"
)

// TagPO 是 tags 的数据库存储模型（Persistence Object）
type TagPO struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement"`
	TagID string `gorm:"column:tag_id;type:uuid;not null;uniqueIndex"`
	Name  string `gorm:"column:name;type:varchar(64);not null;uniqueIndex"`
}

func (TagPO) TableName() string { return "tags" }

// TagPOFromDomain 从领域模型转换为 PO
func TagPOFromDomain(e *domain.Tag) *TagPO {
	if e == nil {
		return nil
	}
	return &TagPO{
		ID:    e.ID,
		TagID: string(e.TagID),
		Name:  e.Name,
	}
}

// ToDomain 转换为领域模型
func (p *TagPO) ToDomain() *domain.Tag {
	if p == nil {
		return nil
	}
	return &domain.Tag{
		ID:    p.ID,
		TagID: domain.TagID(p.TagID),
		Name:  p.Name,
	}
}

// KnowledgeNodeTagPO 是 knowledge_node_tags 的数据库存储模型（Persistence Object）
type KnowledgeNodeTagPO struct {
	NodeID string `gorm:"column:node_id;type:uuid;primaryKey;index:idx_node_tags_node"`
	TagID  string `gorm:"column:tag_id;type:uuid;primaryKey;index:idx_node_tags_tag"`
}

func (KnowledgeNodeTagPO) TableName() string { return "knowledge_node_tags" }

// KnowledgeNodeTagPOFromDomain 从领域模型转换为 PO
func KnowledgeNodeTagPOFromDomain(e *domain.KnowledgeNodeTag) *KnowledgeNodeTagPO {
	if e == nil {
		return nil
	}
	return &KnowledgeNodeTagPO{
		NodeID: string(e.NodeID),
		TagID:  string(e.TagID),
	}
}

// ToDomain 转换为领域模型
func (p *KnowledgeNodeTagPO) ToDomain() *domain.KnowledgeNodeTag {
	if p == nil {
		return nil
	}
	return &domain.KnowledgeNodeTag{
		NodeID: domain.KnowledgeID(p.NodeID),
		TagID:  domain.TagID(p.TagID),
	}
}
