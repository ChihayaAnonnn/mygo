package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"mygo/internal/infra"
	"mygo/internal/user/domain"
)

const (
	sessionKeyPrefix = "session:"
	sessionTTL       = 24 * time.Hour
)

// SessionCache 会话缓存实现
type SessionCache struct {
	redis *infra.RedisClient
}

// NewSessionCache 构造函数
func NewSessionCache(res *infra.Resources) (*SessionCache, error) {
	if res == nil {
		return nil, errors.New("session cache: resources is nil")
	}
	if res.Redis == nil {
		return nil, errors.New("session cache: redis is nil")
	}
	return &SessionCache{redis: res.Redis}, nil
}

// Set 存储会话
func (c *SessionCache) Set(ctx context.Context, sessionID string, data *domain.SessionData) error {
	if c.redis == nil {
		return errors.New("session cache: redis is nil")
	}

	key := sessionKeyPrefix + sessionID
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("session cache: marshal error: %w", err)
	}

	return c.redis.Set(ctx, key, jsonData, sessionTTL).Err()
}

// Get 获取会话
func (c *SessionCache) Get(ctx context.Context, sessionID string) (*domain.SessionData, error) {
	if c.redis == nil {
		return nil, errors.New("session cache: redis is nil")
	}

	key := sessionKeyPrefix + sessionID
	val, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var data domain.SessionData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, fmt.Errorf("session cache: unmarshal error: %w", err)
	}

	return &data, nil
}

// Delete 删除会话
func (c *SessionCache) Delete(ctx context.Context, sessionID string) error {
	if c.redis == nil {
		return errors.New("session cache: redis is nil")
	}

	key := sessionKeyPrefix + sessionID
	return c.redis.Del(ctx, key).Err()
}

// Refresh 刷新会话过期时间
func (c *SessionCache) Refresh(ctx context.Context, sessionID string) error {
	if c.redis == nil {
		return errors.New("session cache: redis is nil")
	}

	key := sessionKeyPrefix + sessionID
	return c.redis.Expire(ctx, key, sessionTTL).Err()
}

// 确保 SessionCache 实现了 domain.SessionCache 接口
var _ domain.SessionCache = (*SessionCache)(nil)
