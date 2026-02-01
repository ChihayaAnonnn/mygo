package config

import (
	"os"

	"mygo/internal/infra"
)

// Config 应用配置
type Config struct {
	Server ServerConfig
	Infra  infra.Config
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string // gin mode: debug, release, test
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Infra: infra.Config{
			PGDSN:    getEnv("DATABASE_URL", "postgres://mygo:avemujica@localhost:5432/app?sslmode=disable"),
			RedisURL: getEnv("REDIS_URL", "redis://localhost:6379/0"),
		},
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
