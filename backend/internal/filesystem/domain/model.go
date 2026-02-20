package domain

import "time"

// ObjectMeta 对象元数据（不包含 payload bytes）。
type ObjectMeta struct {
	Namespace Namespace
	Key       Key

	Backend     string
	ContentType string
	SizeBytes   int64

	ETag     string
	Checksum string
	Metadata map[string]any

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

