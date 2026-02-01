package filesystem

// FileStorage 文件存储接口（预留）
// 用于存储知识相关的文件（如图片、附件等）
type FileStorage interface {
	// Upload 上传文件
	// Upload(ctx context.Context, path string, content []byte) error

	// Download 下载文件
	// Download(ctx context.Context, path string) ([]byte, error)

	// Delete 删除文件
	// Delete(ctx context.Context, path string) error
}

// LocalStorage 本地文件存储实现（预留）
type LocalStorage struct {
	basePath string
}

// NewLocalStorage 构造函数
func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{basePath: basePath}
}
