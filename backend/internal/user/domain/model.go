package domain

// User 用户领域模型
type User struct {
	ID       int64
	UserID   int64
	Username string
	Email    string
	Password string
	Avatar   string
}
