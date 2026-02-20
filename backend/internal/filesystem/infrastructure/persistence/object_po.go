package persistence

import "gorm.io/gorm"

// ObjectPO 是 fs_objects 的数据库存储模型（Persistence Object）
//
// 说明：
// - 该表只存对象元数据，不存 payload（bytes）
// - 不做对象版本化：同一 (namespace, key) 只保留一条记录
type ObjectPO struct {
	gorm.Model

	Namespace string `gorm:"column:namespace;type:varchar(64);not null;uniqueIndex:uk_fs_objects_namespace_key;index:idx_fs_objects_namespace"`
	Key       string `gorm:"column:key;type:text;not null;uniqueIndex:uk_fs_objects_namespace_key"`

	Backend     string `gorm:"column:backend;type:varchar(16);not null;default:'local'"`
	ContentType string `gorm:"column:content_type;type:varchar(128)"`
	SizeBytes   int64  `gorm:"column:size_bytes;type:bigint;not null;default:0"`

	ETag     string `gorm:"column:etag;type:varchar(128)"`
	Checksum string `gorm:"column:checksum;type:varchar(128)"`

	Metadata map[string]any `gorm:"column:metadata;type:jsonb"`
}

func (ObjectPO) TableName() string { return "fs_objects" }
