package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mygo/internal/server"
)

// HTTPServerConfig HTTP 服务器配置
type HTTPServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DefaultHTTPServerConfig 返回默认配置
func DefaultHTTPServerConfig(port string) HTTPServerConfig {
	return HTTPServerConfig{
		Port:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// RunHTTPServer 启动 HTTP 服务器（阻塞，支持优雅关闭）
func RunHTTPServer(app *App) error {
	cfg := DefaultHTTPServerConfig(app.Config.Server.Port)
	return RunHTTPServerWithConfig(app, cfg)
}

// RunHTTPServerWithConfig 使用自定义配置启动 HTTP 服务器
func RunHTTPServerWithConfig(app *App, cfg HTTPServerConfig) error {
	// 创建路由
	router := server.NewRouter(app.RouterConfig())
	log.Println("Router initialized")

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// 启动服务器（非阻塞）
	go func() {
		log.Printf("HTTP server starting on port %s...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down HTTP server...")

	// 优雅关闭（5 秒超时）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("HTTP server stopped gracefully")
	return nil
}
