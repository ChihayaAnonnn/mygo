package po

import (
	"mygo/internal/models/entity"
)

// TagPO 是 tags 的数据库存储模型（Persistence Object）
type TagPO struct {
	ID   string `gorm:"column:id;type:uuid;primaryKey"`
	Name string `gorm:"column:name;type:varchar(64);not null;uniqueIndex"`
}

func (TagPO) TableName() string { return "tags" }

func TagPOFromEntity(e *entity.Tag) *TagPO {
	if e == nil {
		return nil
	}
	return &TagPO{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (p *TagPO) ToEntity() *entity.Tag {
	if p == nil {
		return nil
	}
	return &entity.Tag{
		ID:   p.ID,
		Name: p.Name,
	}
}

// KnowledgeNodeTagPO 是 knowledge_node_tags 的数据库存储模型（Persistence Object）
type KnowledgeNodeTagPO struct {
	NodeID string `gorm:"column:node_id;type:uuid;primaryKey"`
	TagID  string `gorm:"column:tag_id;type:uuid;primaryKey"`
}

func (KnowledgeNodeTagPO) TableName() string { return "knowledge_node_tags" }

func KnowledgeNodeTagPOFromEntity(e *entity.KnowledgeNodeTag) *KnowledgeNodeTagPO {
	if e == nil {
		return nil
	}
	return &KnowledgeNodeTagPO{
		NodeID: e.NodeID,
		TagID:  e.TagID,
	}
}

func (p *KnowledgeNodeTagPO) ToEntity() *entity.KnowledgeNodeTag {
	if p == nil {
		return nil
	}
	return &entity.KnowledgeNodeTag{
		NodeID: p.NodeID,
		TagID:  p.TagID,
	}
}
