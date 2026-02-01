package bootstrap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// WorkerConfig 后台任务配置
type WorkerConfig struct {
	// 并发 Worker 数量
	Concurrency int
}

// DefaultWorkerConfig 返回默认配置
func DefaultWorkerConfig() WorkerConfig {
	return WorkerConfig{
		Concurrency: 4,
	}
}

// RunWorker 启动后台任务处理器（阻塞，支持优雅关闭）
func RunWorker(app *App) error {
	cfg := DefaultWorkerConfig()
	return RunWorkerWithConfig(app, cfg)
}

// RunWorkerWithConfig 使用自定义配置启动后台任务处理器
func RunWorkerWithConfig(app *App, cfg WorkerConfig) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("Worker starting with concurrency=%d...", cfg.Concurrency)

	// TODO: 启动后台任务处理
	// 示例任务类型：
	// - AI Embedding 生成
	// - Markdown 渲染
	// - 知识索引重建
	// - 定时清理任务

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Shutting down worker...")
		cancel()
	case <-ctx.Done():
	}

	// TODO: 等待所有任务完成

	log.Println("Worker stopped gracefully")
	return nil
}
