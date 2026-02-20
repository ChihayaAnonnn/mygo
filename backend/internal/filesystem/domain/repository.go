package domain

import (
	"context"
	"io"
)

// ObjectRepository 管理对象元数据（fs_objects）。
//
// 约束：
// - 删除为软删除（DeletedAt），不做物理清理
// - 查询默认不返回已删除对象
type ObjectRepository interface {
	UpsertMeta(ctx context.Context, meta *ObjectMeta) error
	Head(ctx context.Context, namespace Namespace, key Key) (*ObjectMeta, error)
	ListByPrefix(ctx context.Context, namespace Namespace, prefix string, limit, offset int) ([]*ObjectMeta, error)

	SoftDelete(ctx context.Context, namespace Namespace, key Key) error
	Undelete(ctx context.Context, namespace Namespace, key Key) error
}

type PutResult struct {
	SizeBytes int64
	Checksum  string
	ETag      string
}

// BlobStore 管理对象 payload（bytes）。
//
// 注意：BlobStore 不负责元数据登记与软删除语义。
type BlobStore interface {
	Put(ctx context.Context, namespace Namespace, key Key, payload []byte) (*PutResult, error)
	Get(ctx context.Context, namespace Namespace, key Key) (io.ReadCloser, error)
}
