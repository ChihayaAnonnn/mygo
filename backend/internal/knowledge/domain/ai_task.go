package domain

import (
	"errors"
	"time"
)

// AITaskID AI 任务的唯一标识
type AITaskID string

// AITask AI 处理流程的可追溯任务记录
type AITask struct {
	ID        AITaskID     // AI 任务唯一标识（UUID）
	NodeID    KnowledgeID  // 关联的知识节点
	Version   *int         // 任务对应的 Markdown 版本
	TaskType  AITaskType   // 任务类型
	Status    AITaskStatus // 任务状态
	CreatedAt time.Time    // 任务创建时间
	UpdatedAt time.Time    // 最近一次状态更新
}

// ==================== 状态机方法 ====================

// Start 开始执行任务
func (t *AITask) Start() error {
	if !t.Status.CanTransitionTo(AITaskStatusRunning) {
		return errors.New("cannot start task: invalid status transition")
	}
	t.Status = AITaskStatusRunning
	t.UpdatedAt = time.Now()
	return nil
}

// Complete 标记任务完成
func (t *AITask) Complete() error {
	if !t.Status.CanTransitionTo(AITaskStatusDone) {
		return errors.New("cannot complete task: invalid status transition")
	}
	t.Status = AITaskStatusDone
	t.UpdatedAt = time.Now()
	return nil
}

// Fail 标记任务失败
func (t *AITask) Fail() error {
	if !t.Status.CanTransitionTo(AITaskStatusFailed) {
		return errors.New("cannot fail task: invalid status transition")
	}
	t.Status = AITaskStatusFailed
	t.UpdatedAt = time.Now()
	return nil
}

// Retry 重试失败的任务
func (t *AITask) Retry() error {
	if !t.Status.CanTransitionTo(AITaskStatusPending) {
		return errors.New("cannot retry task: invalid status transition")
	}
	t.Status = AITaskStatusPending
	t.UpdatedAt = time.Now()
	return nil
}

// IsPending 是否等待中
func (t *AITask) IsPending() bool {
	return t.Status == AITaskStatusPending
}

// IsRunning 是否运行中
func (t *AITask) IsRunning() bool {
	return t.Status == AITaskStatusRunning
}

// IsDone 是否已完成
func (t *AITask) IsDone() bool {
	return t.Status == AITaskStatusDone
}

// IsFailed 是否已失败
func (t *AITask) IsFailed() bool {
	return t.Status == AITaskStatusFailed
}
