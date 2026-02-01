package application

import (
	"context"

	"mygo/internal/knowledge/domain"
)

// ==================== KnowledgeApplicationService（应用层编排） ====================

// KnowledgeApplicationService 应用层编排接口
// 职责：编排多个底层 Service，对外提供高层用例（Use Case）
type KnowledgeApplicationService interface {
	// PublishKnowledge 发布知识
	// 流程：更新状态 -> 渲染文件 -> 构建 Chunk -> 生成 Embedding
	PublishKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) error

	// RebuildIndex 重建知识索引
	// 流程：删除旧 Chunk/Embedding -> 重新构建 Chunk -> 重新生成 Embedding
	RebuildIndex(ctx context.Context, knowledgeID domain.KnowledgeID) error
}

// ==================== AppService 实现 ====================

// AppService 应用服务实现
type AppService struct {
	// 领域服务依赖
	knowledgeSvc domain.KnowledgeService
	versionSvc   domain.KnowledgeVersionService
	renderSvc    domain.MarkdownRenderService
	chunkSvc     domain.KnowledgeChunkService
	embeddingSvc domain.EmbeddingService
	retrievalSvc domain.RetrievalService

	// 仓储依赖（用于事务操作）
	knowledgeRepo domain.KnowledgeRepository
	versionRepo   domain.VersionRepository
	chunkRepo     domain.ChunkRepository
	embeddingRepo domain.EmbeddingRepository
}

// NewAppService 构造函数
func NewAppService(
	knowledgeSvc domain.KnowledgeService,
	versionSvc domain.KnowledgeVersionService,
	renderSvc domain.MarkdownRenderService,
	chunkSvc domain.KnowledgeChunkService,
	embeddingSvc domain.EmbeddingService,
	retrievalSvc domain.RetrievalService,
	knowledgeRepo domain.KnowledgeRepository,
	versionRepo domain.VersionRepository,
	chunkRepo domain.ChunkRepository,
	embeddingRepo domain.EmbeddingRepository,
) *AppService {
	return &AppService{
		knowledgeSvc:  knowledgeSvc,
		versionSvc:    versionSvc,
		renderSvc:     renderSvc,
		chunkSvc:      chunkSvc,
		embeddingSvc:  embeddingSvc,
		retrievalSvc:  retrievalSvc,
		knowledgeRepo: knowledgeRepo,
		versionRepo:   versionRepo,
		chunkRepo:     chunkRepo,
		embeddingRepo: embeddingRepo,
	}
}

// PublishKnowledge 发布知识
func (s *AppService) PublishKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) error {
	// TODO: 实现发布流程
	// 1. 获取知识及其最新版本
	// 2. 更新知识状态为 published
	// 3. 渲染 Markdown 到文件
	// 4. 构建 Chunk
	// 5. 生成 Embedding
	return nil
}

// RebuildIndex 重建知识索引
func (s *AppService) RebuildIndex(ctx context.Context, knowledgeID domain.KnowledgeID) error {
	// TODO: 实现重建索引流程
	// 1. 获取知识的最新版本
	// 2. 删除旧的 Chunk 和 Embedding
	// 3. 重新构建 Chunk
	// 4. 重新生成 Embedding
	return nil
}

// 确保 AppService 实现了 KnowledgeApplicationService 接口
var _ KnowledgeApplicationService = (*AppService)(nil)
