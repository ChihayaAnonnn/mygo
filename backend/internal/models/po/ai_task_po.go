package po

import (
	"time"

	"mygo/internal/models/entity"
)

// AITaskPO 是 ai_tasks 的数据库存储模型（Persistence Object）
type AITaskPO struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey"`
	NodeID    string    `gorm:"column:node_id;type:uuid"`
	Version   *int      `gorm:"column:version;type:int"`
	TaskType  string    `gorm:"column:task_type;type:varchar(32)"`
	Status    string    `gorm:"column:status;type:varchar(16)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (AITaskPO) TableName() string { return "ai_tasks" }

func AITaskPOFromEntity(e *entity.AITask) *AITaskPO {
	if e == nil {
		return nil
	}
	return &AITaskPO{
		ID:        e.ID,
		NodeID:    e.NodeID,
		Version:   e.Version,
		TaskType:  string(e.TaskType),
		Status:    string(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (p *AITaskPO) ToEntity() *entity.AITask {
	if p == nil {
		return nil
	}
	return &entity.AITask{
		ID:        p.ID,
		NodeID:    p.NodeID,
		Version:   p.Version,
		TaskType:  entity.AITaskType(p.TaskType),
		Status:    entity.AITaskStatus(p.Status),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
