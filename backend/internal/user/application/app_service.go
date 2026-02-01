package application

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"mygo/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

// AppService 用户应用服务（用例层实现）
// 负责编排领域对象完成业务用例
type AppService struct {
	userRepo     domain.UserRepository
	sessionCache domain.SessionCache
}

// NewAppService 构造函数
func NewAppService(userRepo domain.UserRepository, sessionCache domain.SessionCache) *AppService {
	return &AppService{
		userRepo:     userRepo,
		sessionCache: sessionCache,
	}
}

// Register 用户注册
func (s *AppService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, domain.ErrInvalidInput
	}

	// 检查用户名是否已存在
	existing, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, fmt.Errorf("check username: %w", err)
	}
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// 检查邮箱是否已存在
	existing, err = s.userRepo.GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, fmt.Errorf("check email: %w", err)
	}
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	// 生成用户 ID (简单实现，生产环境应使用雪花算法等)
	userID := time.Now().UnixNano()

	user := &domain.User{
		UserID:   userID,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

// Login 用户登录
func (s *AppService) Login(ctx context.Context, username, password string) (string, *domain.User, error) {
	if username == "" || password == "" {
		return "", nil, domain.ErrInvalidInput
	}

	// 查找用户
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", nil, domain.ErrInvalidCredentials
		}
		return "", nil, fmt.Errorf("get user: %w", err)
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, domain.ErrInvalidCredentials
	}

	// 生成会话 ID
	sessionID, err := generateSessionID()
	if err != nil {
		return "", nil, fmt.Errorf("generate session: %w", err)
	}

	// 存储会话到缓存
	if s.sessionCache != nil {
		sessionData := &domain.SessionData{
			UserID:    user.UserID,
			Username:  user.Username,
			CreatedAt: time.Now().Unix(),
		}
		if err := s.sessionCache.Set(ctx, sessionID, sessionData); err != nil {
			return "", nil, fmt.Errorf("save session: %w", err)
		}
	}

	return sessionID, user, nil
}

// Logout 用户登出
func (s *AppService) Logout(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return domain.ErrInvalidInput
	}

	if s.sessionCache != nil {
		return s.sessionCache.Delete(ctx, sessionID)
	}
	return nil
}

// GetUserByID 根据 ID 获取用户
func (s *AppService) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	if id == 0 {
		return nil, domain.ErrInvalidInput
	}

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 清除敏感信息
	user.Password = ""
	return user, nil
}

// generateSessionID 生成会话 ID
func generateSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// 确保 AppService 实现了 domain.UserService 接口
var _ domain.UserService = (*AppService)(nil)
