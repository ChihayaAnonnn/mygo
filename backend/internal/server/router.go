package server

import (
	knowledgeHttp "mygo/internal/knowledge/interfaces/http"
	userHttp "mygo/internal/user/interfaces/http"

	"github.com/gin-gonic/gin"
)

// RouterConfig 路由配置
type RouterConfig struct {
	UserHandler      *userHttp.Handler
	KnowledgeHandler *knowledgeHttp.Handler
}

// NewRouter 创建路由
func NewRouter(cfg RouterConfig) *gin.Engine {
	r := gin.New()

	// 全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok":      true,
			"message": "service is healthy",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// 注册各领域模块路由
		if cfg.UserHandler != nil {
			userHttp.RegisterRoutes(api, cfg.UserHandler)
		}
		if cfg.KnowledgeHandler != nil {
			knowledgeHttp.RegisterRoutes(api, cfg.KnowledgeHandler)
		}
	}

	// 兼容旧路由 (IM 服务)
	im := r.Group("/im")
	{
		im.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"ok":   true,
				"msg":  "Go service is running",
				"path": "/im/health",
			})
		})

		// WebSocket 预留
		im.GET("/ws", func(c *gin.Context) {
			c.String(200, "WebSocket endpoint placeholder")
		})
	}

	return r
}
