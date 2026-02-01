package http

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 用户登录响应
type LoginResponse struct {
	SessionID string `json:"session_id"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
}
