package application

import (
	"context"
	"fmt"
	"time"

	"mygo/internal/knowledge/domain"

	"github.com/google/uuid"
)

// knowledgeServiceImpl 知识元服务实现
type knowledgeServiceImpl struct {
	repo domain.KnowledgeRepository
}

// NewKnowledgeService 创建知识元服务
func NewKnowledgeService(repo domain.KnowledgeRepository) domain.KnowledgeService {
	return &knowledgeServiceImpl{repo: repo}
}

// CreateKnowledge 创建知识
func (s *knowledgeServiceImpl) CreateKnowledge(ctx context.Context, cmd domain.CreateKnowledgeCmd) (domain.KnowledgeID, error) {
	if cmd.Title == "" {
		return "", fmt.Errorf("%w: title is required", domain.ErrInvalidInput)
	}
	if cmd.NodeType == "" {
		return "", fmt.Errorf("%w: node_type is required", domain.ErrInvalidInput)
	}

	nodeID := domain.KnowledgeID(uuid.New().String())
	now := time.Now()

	node := &domain.Node{
		NodeID:         string(nodeID),
		NodeType:       cmd.NodeType,
		Title:          cmd.Title,
		Summary:        cmd.Summary,
		Status:         domain.NodeStatusDraft,
		CurrentVersion: 0,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.Create(ctx, node); err != nil {
		return "", fmt.Errorf("create knowledge: %w", err)
	}

	return domain.KnowledgeID(node.NodeID), nil
}

// UpdateKnowledgeMeta 更新知识元信息
func (s *knowledgeServiceImpl) UpdateKnowledgeMeta(ctx context.Context, cmd domain.UpdateKnowledgeMetaCmd) error {
	if cmd.ID == "" {
		return fmt.Errorf("%w: knowledge id is required", domain.ErrInvalidInput)
	}

	node, err := s.repo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if cmd.Title != nil {
		node.Title = *cmd.Title
	}
	if cmd.Summary != nil {
		node.Summary = *cmd.Summary
	}
	if cmd.Status != nil {
		node.Status = *cmd.Status
	}
	if cmd.Confidence != nil {
		node.Confidence = cmd.Confidence
	}
	node.UpdatedAt = time.Now()

	return s.repo.Update(ctx, node)
}

// GetKnowledge 获取知识
func (s *knowledgeServiceImpl) GetKnowledge(ctx context.Context, id domain.KnowledgeID) (*domain.Node, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: knowledge id is required", domain.ErrInvalidInput)
	}
	return s.repo.GetByID(ctx, id)
}

// ListKnowledge 列出知识
func (s *knowledgeServiceImpl) ListKnowledge(ctx context.Context, query domain.KnowledgeQuery) ([]*domain.Node, error) {
	return s.repo.List(ctx, query)
}

// ArchiveKnowledge 归档知识（软删除）
func (s *knowledgeServiceImpl) ArchiveKnowledge(ctx context.Context, id domain.KnowledgeID) error {
	if id == "" {
		return fmt.Errorf("%w: knowledge id is required", domain.ErrInvalidInput)
	}

	node, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	node.Status = domain.NodeStatusArchived
	node.UpdatedAt = time.Now()

	return s.repo.Update(ctx, node)
}

// 编译时检查
var _ domain.KnowledgeService = (*knowledgeServiceImpl)(nil)
