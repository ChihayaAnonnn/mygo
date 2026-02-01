package domain

import "context"

// UserService 用户领域服务接口（用例层实现）
type UserService interface {
	// Register 用户注册
	Register(ctx context.Context, username, email, password string) (*User, error)

	// Login 用户登录，返回 sessionID
	Login(ctx context.Context, username, password string) (sessionID string, user *User, err error)

	// Logout 用户登出
	Logout(ctx context.Context, sessionID string) error

	// GetUserByID 根据 ID 获取用户（不含敏感信息）
	GetUserByID(ctx context.Context, id int64) (*User, error)
}
