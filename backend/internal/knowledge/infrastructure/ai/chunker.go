package ai

import "context"

// ChunkResult 分块结果
type ChunkResult struct {
	Content     string // 分块内容
	HeadingPath string // 标题路径
	TokenCount  int    // Token 数量
	ChunkIndex  int    // 分块索引
}

// Chunker 文本分块接口（预留）
// 用于将 Markdown 内容分割为适合 AI 处理的块
type Chunker interface {
	// Chunk 将 Markdown 内容分块
	Chunk(ctx context.Context, markdown string) ([]*ChunkResult, error)

	// MaxTokens 返回每个块的最大 Token 数
	MaxTokens() int
}

// MarkdownChunker Markdown 分块实现（预留）
type MarkdownChunker struct {
	maxTokens int
	overlap   int
}

// NewMarkdownChunker 构造函数
func NewMarkdownChunker(maxTokens, overlap int) *MarkdownChunker {
	return &MarkdownChunker{
		maxTokens: maxTokens,
		overlap:   overlap,
	}
}

// MaxTokens 返回最大 Token 数
func (c *MarkdownChunker) MaxTokens() int {
	return c.maxTokens
}
