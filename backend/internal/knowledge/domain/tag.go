package domain

import "errors"

// ==================== Tag 错误 ====================

var ErrTagNotFound = errors.New("tag not found")

// ==================== Tag ID ====================

// TagID 标签的唯一标识
type TagID string

// ==================== Tag 模型 ====================

// Tag 弱语义分类标签
type Tag struct {
	ID    int64  // 数据库自增主键（内部使用）
	TagID TagID  // 标签唯一标识（UUID，对外暴露）
	Name  string // 标签名称（唯一）
}

// ==================== KnowledgeNodeTag 关联 ====================

// KnowledgeNodeTag 知识节点与标签的关联
type KnowledgeNodeTag struct {
	NodeID KnowledgeID // 知识节点 ID
	TagID  TagID       // 标签 ID
}
