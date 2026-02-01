package entity

import (
	"context"
	"errors"
)

var ErrTagNotFound = errors.New("tag not found")

// Tag 弱语义分类标签
type Tag struct {
	ID   string // 标签唯一标识（UUID）
	Name string // 标签名称（唯一）
}

// TagRepository 定义标签的仓储接口
type TagRepository interface {
	Create(ctx context.Context, tag *Tag) error
	GetByID(ctx context.Context, id string) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Tag, error)
	// GetOrCreate 获取或创建标签（幂等操作）
	GetOrCreate(ctx context.Context, name string) (*Tag, error)
}

// KnowledgeNodeTag 知识节点与标签的关联
type KnowledgeNodeTag struct {
	NodeID string // 知识节点 ID
	TagID  string // 标签 ID
}

// KnowledgeNodeTagRepository 定义节点标签关联的仓储接口
type KnowledgeNodeTagRepository interface {
	Create(ctx context.Context, nodeTag *KnowledgeNodeTag) error
	Delete(ctx context.Context, nodeID, tagID string) error
	ListTagsByNode(ctx context.Context, nodeID string) ([]*Tag, error)
	ListNodesByTag(ctx context.Context, tagID string) ([]string, error)
	SetNodeTags(ctx context.Context, nodeID string, tagIDs []string) error
}
