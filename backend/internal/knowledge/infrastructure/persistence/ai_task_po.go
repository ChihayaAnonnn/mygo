package persistence

import (
	"time"

	"mygo/internal/knowledge/domain"
)

// AITaskPO 是 ai_tasks 的数据库存储模型（Persistence Object）
type AITaskPO struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey"`
	NodeID    string    `gorm:"column:node_id;type:uuid;index:idx_ai_tasks_node"`
	Version   *int      `gorm:"column:version;type:int"`
	TaskType  string    `gorm:"column:task_type;type:varchar(32);not null;index:idx_ai_tasks_type"`
	Status    string    `gorm:"column:status;type:varchar(16);not null;index:idx_ai_tasks_status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (AITaskPO) TableName() string { return "ai_tasks" }

// AITaskPOFromDomain 从领域模型转换为 PO
func AITaskPOFromDomain(e *domain.AITask) *AITaskPO {
	if e == nil {
		return nil
	}
	return &AITaskPO{
		ID:        string(e.ID),
		NodeID:    string(e.NodeID),
		Version:   e.Version,
		TaskType:  string(e.TaskType),
		Status:    string(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// ToDomain 转换为领域模型
func (p *AITaskPO) ToDomain() *domain.AITask {
	if p == nil {
		return nil
	}
	return &domain.AITask{
		ID:        domain.AITaskID(p.ID),
		NodeID:    domain.KnowledgeID(p.NodeID),
		Version:   p.Version,
		TaskType:  domain.AITaskType(p.TaskType),
		Status:    domain.AITaskStatus(p.Status),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
