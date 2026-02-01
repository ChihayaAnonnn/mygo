package main

import (
	"log"

	"mygo/internal/bootstrap"
)

func main() {
	// 1. 初始化应用
	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer func() {
		if err := app.Close(); err != nil {
			log.Printf("Error closing app: %v", err)
		}
	}()

	// 2. 启动 HTTP 服务器（阻塞）
	if err := bootstrap.RunHTTPServer(app); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
