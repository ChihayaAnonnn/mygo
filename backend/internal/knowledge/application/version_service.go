package application

import (
	"context"
	"fmt"
	"time"

	"mygo/internal/knowledge/domain"

	"github.com/google/uuid"
)

// versionServiceImpl 版本服务实现
type versionServiceImpl struct {
	versionRepo   domain.VersionRepository
	knowledgeRepo domain.KnowledgeRepository
}

// NewVersionService 创建版本服务
func NewVersionService(
	versionRepo domain.VersionRepository,
	knowledgeRepo domain.KnowledgeRepository,
) domain.KnowledgeVersionService {
	return &versionServiceImpl{
		versionRepo:   versionRepo,
		knowledgeRepo: knowledgeRepo,
	}
}

// CreateVersion 创建新版本
func (s *versionServiceImpl) CreateVersion(ctx context.Context, cmd domain.CreateVersionCmd) (domain.VersionID, error) {
	if cmd.KnowledgeID == "" {
		return "", fmt.Errorf("%w: knowledge_id is required", domain.ErrInvalidInput)
	}
	if cmd.ContentMd == "" {
		return "", fmt.Errorf("%w: content_md is required", domain.ErrInvalidInput)
	}

	// 验证知识节点存在
	node, err := s.knowledgeRepo.GetByID(ctx, cmd.KnowledgeID)
	if err != nil {
		return "", err
	}

	// 计算下一个版本号
	nextVersion := node.CurrentVersion + 1

	versionID := domain.VersionID(uuid.New().String())
	version := &domain.Version{
		VersionID: string(versionID),
		NodeID:    string(cmd.KnowledgeID),
		Version:   nextVersion,
		ContentMd: cmd.ContentMd,
		CreatedAt: time.Now(),
	}

	if err := s.versionRepo.Create(ctx, version); err != nil {
		return "", fmt.Errorf("create version: %w", err)
	}

	// 更新知识节点的当前版本号
	node.CurrentVersion = nextVersion
	if err := s.knowledgeRepo.Update(ctx, node); err != nil {
		return "", fmt.Errorf("update current version: %w", err)
	}

	return domain.VersionID(version.VersionID), nil
}

// GetVersion 获取指定版本
func (s *versionServiceImpl) GetVersion(ctx context.Context, id domain.VersionID) (*domain.Version, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: version_id is required", domain.ErrInvalidInput)
	}
	return s.versionRepo.GetByID(ctx, id)
}

// GetLatestVersion 获取知识的最新版本
func (s *versionServiceImpl) GetLatestVersion(ctx context.Context, knowledgeID domain.KnowledgeID) (*domain.Version, error) {
	if knowledgeID == "" {
		return nil, fmt.Errorf("%w: knowledge_id is required", domain.ErrInvalidInput)
	}
	return s.versionRepo.GetLatestByKnowledge(ctx, knowledgeID)
}

// ListVersions 列出知识的所有版本
func (s *versionServiceImpl) ListVersions(ctx context.Context, knowledgeID domain.KnowledgeID) ([]*domain.Version, error) {
	if knowledgeID == "" {
		return nil, fmt.Errorf("%w: knowledge_id is required", domain.ErrInvalidInput)
	}
	return s.versionRepo.ListByKnowledge(ctx, knowledgeID)
}

// 编译时检查
var _ domain.KnowledgeVersionService = (*versionServiceImpl)(nil)
