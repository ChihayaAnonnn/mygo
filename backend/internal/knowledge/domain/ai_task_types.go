package domain

import "errors"

// ==================== AI Task 错误 ====================

var ErrAITaskNotFound = errors.New("ai task not found")

// ==================== AI Task 类型 ====================

// AITaskType 定义 AI 任务类型
type AITaskType string

const (
	AITaskTypeChunk     AITaskType = "chunk"     // 文本分块
	AITaskTypeEmbedding AITaskType = "embedding" // 向量生成
	AITaskTypeSummary   AITaskType = "summary"   // 摘要生成
	AITaskTypeEdge      AITaskType = "edge"      // 关系推断
)

// ==================== AI Task 状态 ====================

// AITaskStatus 定义 AI 任务状态
type AITaskStatus string

const (
	AITaskStatusPending AITaskStatus = "pending" // 等待处理
	AITaskStatusRunning AITaskStatus = "running" // 处理中
	AITaskStatusDone    AITaskStatus = "done"    // 完成
	AITaskStatusFailed  AITaskStatus = "failed"  // 失败
)

// ==================== AI Task 状态机 ====================

// ValidTransitions 定义合法的状态转换
var ValidTransitions = map[AITaskStatus][]AITaskStatus{
	AITaskStatusPending: {AITaskStatusRunning, AITaskStatusFailed},
	AITaskStatusRunning: {AITaskStatusDone, AITaskStatusFailed},
	AITaskStatusFailed:  {AITaskStatusPending}, // 重试
	AITaskStatusDone:    {},                    // 终态
}

// CanTransitionTo 检查是否可以转换到目标状态
func (s AITaskStatus) CanTransitionTo(target AITaskStatus) bool {
	allowed, ok := ValidTransitions[s]
	if !ok {
		return false
	}
	for _, t := range allowed {
		if t == target {
			return true
		}
	}
	return false
}
