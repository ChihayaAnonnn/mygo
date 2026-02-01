package ai

import "context"

// Embedder 向量嵌入接口（预留）
// 用于将文本转换为向量表示
type Embedder interface {
	// Embed 将文本转换为向量
	Embed(ctx context.Context, text string) ([]float32, error)

	// BatchEmbed 批量将文本转换为向量
	BatchEmbed(ctx context.Context, texts []string) ([][]float32, error)

	// ModelName 返回使用的模型名称
	ModelName() string

	// Dimension 返回向量维度
	Dimension() int
}

// OpenAIEmbedder OpenAI 嵌入实现（预留）
type OpenAIEmbedder struct {
	apiKey    string
	model     string
	dimension int
}

// NewOpenAIEmbedder 构造函数
func NewOpenAIEmbedder(apiKey, model string) *OpenAIEmbedder {
	return &OpenAIEmbedder{
		apiKey:    apiKey,
		model:     model,
		dimension: 1536, // text-embedding-3-small 默认维度
	}
}

// ModelName 返回模型名称
func (e *OpenAIEmbedder) ModelName() string {
	return e.model
}

// Dimension 返回向量维度
func (e *OpenAIEmbedder) Dimension() int {
	return e.dimension
}
