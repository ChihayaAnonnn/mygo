package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient 作为 infra 暴露给上层的依赖类型别名
type RedisClient = redis.Client

// NewRedis 连接 Redis
// url 示例: redis://:password@localhost:6379/0
func NewRedis(url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	// 覆盖默认配置以优化生产环境
	opts.DialTimeout = 5 * time.Second
	opts.ReadTimeout = 3 * time.Second
	opts.WriteTimeout = 3 * time.Second
	opts.PoolSize = 10
	opts.MinIdleConns = 2

	rdb := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return rdb, nil
}
