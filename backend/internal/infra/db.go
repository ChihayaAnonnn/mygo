package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPG 连接 PostgreSQL
func NewPG(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// 连接池配置
	config.MaxConns = 10                       // 最大连接数
	config.MinConns = 2                        // 最小空闲连接
	config.MaxConnLifetime = 1 * time.Hour     // 连接最大存活时间
	config.MaxConnIdleTime = 30 * time.Minute  // 空闲连接最大存活时间

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Ping 测试
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
