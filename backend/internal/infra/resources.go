package infra

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// RedisClient 作为 infra 暴露给上层的依赖类型别名，
type RedisClient = redis.Client

// Config 统一承载 infra 初始化需要的配置。
type Config struct {
	PGDSN    string
	RedisURL string
}

// Resources 表示应用运行所需的基础设施依赖集合。
// 由 infra 负责创建和释放。
type Resources struct {
	Redis *RedisClient
	DB    *GormDB
}

func NewResources(cfg Config) (*Resources, error) {
	if cfg.PGDSN == "" {
		return nil, fmt.Errorf("infra: PGDSN is empty")
	}
	if cfg.RedisURL == "" {
		return nil, fmt.Errorf("infra: RedisURL is empty")
	}

	rdb, err := NewRedis(cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	gdb, err := NewGormPG(cfg.PGDSN)
	if err != nil {
		_ = rdb.Close()
		return nil, err
	}

	return &Resources{
		Redis: rdb,
		DB:    gdb,
	}, nil
}

func (r *Resources) Close() error {
	if r == nil {
		return nil
	}

	if r.DB != nil {
		if sqlDB, err := r.DB.DB(); err == nil && sqlDB != nil {
			_ = sqlDB.Close()
		}
	}

	if r.Redis != nil {
		return r.Redis.Close()
	}

	return nil
}
