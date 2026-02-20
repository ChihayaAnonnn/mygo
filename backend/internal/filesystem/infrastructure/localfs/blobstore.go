package localfs

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"mygo/internal/filesystem/domain"
)

// BlobStore 本地文件系统 payload 存储实现。
//
// 物理路径：
//   <basePath>/namespaces/<namespace>/<key>
//
// 注意：该实现不处理元数据写入与软删除语义。
type BlobStore struct {
	basePath string
}

func NewLocalFSBlobStore(basePath string) (*BlobStore, error) {
	if strings.TrimSpace(basePath) == "" {
		return nil, errors.New("localfs blobstore: basePath is empty")
	}
	return &BlobStore{basePath: basePath}, nil
}

func (s *BlobStore) Put(ctx context.Context, namespace domain.Namespace, key domain.Key, payload []byte) (*domain.PutResult, error) {
	_ = ctx // 预留：未来可在大文件流式写入时支持 ctx cancel

	if namespace == "" || key == "" {
		return nil, domain.ErrInvalidInput
	}

	target, err := s.objectPath(namespace, key)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	tmp, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return nil, err
	}
	tmpName := tmp.Name()

	cleanup := func() {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
	}

	hasher := sha256.New()
	w := io.MultiWriter(tmp, hasher)
	n, err := w.Write(payload)
	if err != nil {
		cleanup()
		return nil, err
	}
	if n != len(payload) {
		cleanup()
		return nil, io.ErrShortWrite
	}

	if err := tmp.Sync(); err != nil {
		cleanup()
		return nil, err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpName)
		return nil, err
	}

	if err := os.Rename(tmpName, target); err != nil {
		_ = os.Remove(tmpName)
		return nil, err
	}
	_ = os.Chmod(target, 0o644)

	sum := hex.EncodeToString(hasher.Sum(nil))
	return &domain.PutResult{
		SizeBytes: int64(len(payload)),
		Checksum:  "sha256:" + sum,
		ETag:      "sha256:" + sum,
	}, nil
}

func (s *BlobStore) Get(ctx context.Context, namespace domain.Namespace, key domain.Key) (io.ReadCloser, error) {
	_ = ctx
	if namespace == "" || key == "" {
		return nil, domain.ErrInvalidInput
	}
	target, err := s.objectPath(namespace, key)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(target)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}
	return f, nil
}

func (s *BlobStore) objectPath(namespace domain.Namespace, key domain.Key) (string, error) {
	root := filepath.Clean(s.basePath)
	physicalKey := filepath.FromSlash(key.String())
	target := filepath.Join(root, "namespaces", namespace.String(), physicalKey)

	rel, err := filepath.Rel(root, target)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", domain.ErrInvalidInput
	}
	return target, nil
}

var _ domain.BlobStore = (*BlobStore)(nil)

