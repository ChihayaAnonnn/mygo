package domain

import "context"

// UserRepository 用户仓储接口（领域层定义，基础设施层实现）
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}

// SessionData 会话数据
type SessionData struct {
	UserID    int64
	Username  string
	CreatedAt int64 // Unix timestamp
}

// SessionCache 会话缓存接口（领域层定义，基础设施层实现）
type SessionCache interface {
	Set(ctx context.Context, sessionID string, data *SessionData) error
	Get(ctx context.Context, sessionID string) (*SessionData, error)
	Delete(ctx context.Context, sessionID string) error
	Refresh(ctx context.Context, sessionID string) error
}
