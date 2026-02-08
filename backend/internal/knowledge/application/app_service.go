package application

import (
	"context"
	"fmt"
	"time"

	"mygo/internal/knowledge/domain"
)

// ==================== KnowledgeApplicationService（应用层编排） ====================

// KnowledgeApplicationService 应用层编排接口
type KnowledgeApplicationService interface {
	// PublishKnowledge 发布知识（仅更新状态 + 渲染文件）
	// Chunk/Embedding 由 Agent 端主动调用写入 API
	PublishKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) error

	// RebuildIndex 清除旧的 Chunk/Embedding 数据
	// Agent 负责重新生成并写入
	RebuildIndex(ctx context.Context, knowledgeID domain.KnowledgeID) error
}

// ==================== AppService 实现 ====================

// AppService 应用服务实现
type AppService struct {
	knowledgeSvc domain.KnowledgeService
	versionSvc   domain.KnowledgeVersionService
	renderSvc    domain.MarkdownRenderService

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
	knowledgeRepo domain.KnowledgeRepository,
	versionRepo domain.VersionRepository,
	chunkRepo domain.ChunkRepository,
	embeddingRepo domain.EmbeddingRepository,
) *AppService {
	return &AppService{
		knowledgeSvc:  knowledgeSvc,
		versionSvc:    versionSvc,
		renderSvc:     renderSvc,
		knowledgeRepo: knowledgeRepo,
		versionRepo:   versionRepo,
		chunkRepo:     chunkRepo,
		embeddingRepo: embeddingRepo,
	}
}

// PublishKnowledge 发布知识
func (s *AppService) PublishKnowledge(ctx context.Context, knowledgeID domain.KnowledgeID) error {
	// 1. 获取知识节点
	node, err := s.knowledgeRepo.GetByID(ctx, knowledgeID)
	if err != nil {
		return fmt.Errorf("get knowledge: %w", err)
	}

	// 2. 更新状态为 published
	node.Status = domain.NodeStatusPublished
	node.UpdatedAt = time.Now()
	if err := s.knowledgeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("update status: %w", err)
	}

	// 3. 渲染文件（如果 renderSvc 可用）
	if s.renderSvc != nil {
		version, err := s.versionRepo.GetLatestByKnowledge(ctx, knowledgeID)
		if err != nil {
			return fmt.Errorf("get latest version: %w", err)
		}
		if _, err := s.renderSvc.RenderToFile(ctx, version); err != nil {
			return fmt.Errorf("render file: %w", err)
		}
	}

	return nil
}

// RebuildIndex 清除旧的 Chunk/Embedding 数据
func (s *AppService) RebuildIndex(ctx context.Context, knowledgeID domain.KnowledgeID) error {
	// 1. 获取最新版本
	version, err := s.versionRepo.GetLatestByKnowledge(ctx, knowledgeID)
	if err != nil {
		return fmt.Errorf("get latest version: %w", err)
	}

	nodeID := domain.KnowledgeID(version.NodeID)

	// 2. 删除旧的 Embedding
	if s.chunkRepo != nil && s.embeddingRepo != nil {
		oldChunks, err := s.chunkRepo.ListByNodeVersion(ctx, nodeID, version.Version)
		if err == nil && len(oldChunks) > 0 {
			chunkIDs := make([]domain.ChunkID, 0, len(oldChunks))
			for _, c := range oldChunks {
				chunkIDs = append(chunkIDs, domain.ChunkID(c.ChunkID))
			}
			_ = s.embeddingRepo.DeleteByChunkIDs(ctx, chunkIDs)
		}
	}

	// 3. 删除旧的 Chunk
	if s.chunkRepo != nil {
		_ = s.chunkRepo.DeleteByNodeVersion(ctx, nodeID, version.Version)
	}

	return nil
}

// 确保 AppService 实现了 KnowledgeApplicationService 接口
var _ KnowledgeApplicationService = (*AppService)(nil)
