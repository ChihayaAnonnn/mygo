package entity

import (
	"context"
	"errors"
	"time"
)

var ErrAITaskNotFound = errors.New("ai task not found")

// AITaskType 定义 AI 任务类型
type AITaskType string

const (
	AITaskTypeChunk     AITaskType = "chunk"     // 文本分块
	AITaskTypeEmbedding AITaskType = "embedding" // 向量生成
	AITaskTypeSummary   AITaskType = "summary"   // 摘要生成
	AITaskTypeEdge      AITaskType = "edge"      // 关系推断
)

// AITaskStatus 定义 AI 任务状态
type AITaskStatus string

const (
	AITaskStatusPending AITaskStatus = "pending" // 等待处理
	AITaskStatusRunning AITaskStatus = "running" // 处理中
	AITaskStatusDone    AITaskStatus = "done"    // 完成
	AITaskStatusFailed  AITaskStatus = "failed"  // 失败
)

// AITask AI 处理流程的可追溯任务记录
type AITask struct {
	ID        string       // AI 任务唯一标识（UUID）
	NodeID    string       // 关联的知识节点
	Version   *int         // 任务对应的 Markdown 版本
	TaskType  AITaskType   // 任务类型
	Status    AITaskStatus // 任务状态
	CreatedAt time.Time    // 任务创建时间
	UpdatedAt time.Time    // 最近一次状态更新
}

// AITaskRepository 定义 AI 任务的仓储接口
type AITaskRepository interface {
	Create(ctx context.Context, task *AITask) error
	GetByID(ctx context.Context, id string) (*AITask, error)
	Update(ctx context.Context, task *AITask) error
	ListByNode(ctx context.Context, nodeID string) ([]*AITask, error)
	ListByStatus(ctx context.Context, status AITaskStatus) ([]*AITask, error)
	ListPending(ctx context.Context, limit int) ([]*AITask, error)
}
