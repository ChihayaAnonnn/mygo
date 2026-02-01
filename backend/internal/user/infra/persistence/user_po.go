package persistence

import (
	"time"

	"mygo/internal/user/domain"
)

// UserPO 是 user 的数据库存储模型（Persistence Object）。
// 默认映射到表 `users`，字段名使用 snake_case。
type UserPO struct {
	ID     int64 `gorm:"column:id;primaryKey"`
	UserID int64 `gorm:"column:user_id;not null;uniqueIndex"`

	Username string `gorm:"column:username;type:varchar(64);not null;uniqueIndex"`
	Email    string `gorm:"column:email;type:varchar(128);not null;uniqueIndex"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(255)"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (UserPO) TableName() string { return "users" }

// FromDomain 从领域模型转换为 PO
func FromDomain(u *domain.User) *UserPO {
	if u == nil {
		return nil
	}
	return &UserPO{
		ID:       u.ID,
		UserID:   u.UserID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Avatar:   u.Avatar,
	}
}

// ToDomain 转换为领域模型
func (p *UserPO) ToDomain() *domain.User {
	if p == nil {
		return nil
	}
	return &domain.User{
		ID:       p.ID,
		UserID:   p.UserID,
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Avatar:   p.Avatar,
	}
}
