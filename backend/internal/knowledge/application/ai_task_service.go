package application

import (
	"context"
	"time"

	"mygo/internal/knowledge/domain"

	"github.com/google/uuid"
)

// AITaskService AI 任务服务接口
type AITaskService interface {
	// CreateTask 创建新任务
	CreateTask(ctx context.Context, nodeID domain.KnowledgeID, version *int, taskType domain.AITaskType) (*domain.AITask, error)

	// GetTask 获取任务
	GetTask(ctx context.Context, id domain.AITaskID) (*domain.AITask, error)

	// RetryTask 重试失败的任务
	RetryTask(ctx context.Context, id domain.AITaskID) error

	// ListPendingTasks 列出等待处理的任务
	ListPendingTasks(ctx context.Context, limit int) ([]*domain.AITask, error)

	// ListTasksByNode 列出节点的所有任务
	ListTasksByNode(ctx context.Context, nodeID domain.KnowledgeID) ([]*domain.AITask, error)

	// StartTask 开始执行任务
	StartTask(ctx context.Context, id domain.AITaskID) error

	// CompleteTask 标记任务完成
	CompleteTask(ctx context.Context, id domain.AITaskID) error

	// FailTask 标记任务失败
	FailTask(ctx context.Context, id domain.AITaskID) error
}

// aiTaskServiceImpl AI 任务服务实现
type aiTaskServiceImpl struct {
	repo domain.AITaskRepository
}

// NewAITaskService 创建 AI 任务服务
func NewAITaskService(repo domain.AITaskRepository) AITaskService {
	return &aiTaskServiceImpl{repo: repo}
}

// CreateTask 创建新任务
func (s *aiTaskServiceImpl) CreateTask(ctx context.Context, nodeID domain.KnowledgeID, version *int, taskType domain.AITaskType) (*domain.AITask, error) {
	task := &domain.AITask{
		ID:        domain.AITaskID(uuid.New().String()),
		NodeID:    nodeID,
		Version:   version,
		TaskType:  taskType,
		Status:    domain.AITaskStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

// GetTask 获取任务
func (s *aiTaskServiceImpl) GetTask(ctx context.Context, id domain.AITaskID) (*domain.AITask, error) {
	return s.repo.GetByID(ctx, id)
}

// RetryTask 重试失败的任务
func (s *aiTaskServiceImpl) RetryTask(ctx context.Context, id domain.AITaskID) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := task.Retry(); err != nil {
		return err
	}

	return s.repo.Update(ctx, task)
}

// ListPendingTasks 列出等待处理的任务
func (s *aiTaskServiceImpl) ListPendingTasks(ctx context.Context, limit int) ([]*domain.AITask, error) {
	return s.repo.ListPending(ctx, limit)
}

// ListTasksByNode 列出节点的所有任务
func (s *aiTaskServiceImpl) ListTasksByNode(ctx context.Context, nodeID domain.KnowledgeID) ([]*domain.AITask, error) {
	return s.repo.ListByNode(ctx, nodeID)
}

// StartTask 开始执行任务
func (s *aiTaskServiceImpl) StartTask(ctx context.Context, id domain.AITaskID) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := task.Start(); err != nil {
		return err
	}

	return s.repo.Update(ctx, task)
}

// CompleteTask 标记任务完成
func (s *aiTaskServiceImpl) CompleteTask(ctx context.Context, id domain.AITaskID) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := task.Complete(); err != nil {
		return err
	}

	return s.repo.Update(ctx, task)
}

// FailTask 标记任务失败
func (s *aiTaskServiceImpl) FailTask(ctx context.Context, id domain.AITaskID) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := task.Fail(); err != nil {
		return err
	}

	return s.repo.Update(ctx, task)
}
