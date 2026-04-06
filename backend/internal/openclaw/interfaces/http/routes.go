package http

import (
	"github.com/gin-gonic/gin"
	"mygo/backend/internal/openclaw/application"
)

type Router struct {
	service application.OpenClawService
}

func NewRouter(service application.OpenClawService) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) RegisterRoutes(router *gin.RouterGroup) {
	// 动画控制相关路由
	avatar := router.Group("/avatar")
	{
		avatar.POST("/action", r.handleAvatarAction)
		avatar.GET("/actions", r.handleGetAvailableActions)
		avatar.GET("/status", r.handleGetAvatarStatus)
	}

	// 信息管理相关路由
	info := router.Group("/info")
	{
		info.GET("/collections", r.handleGetInfoCollections)
		info.POST("/submit", r.handleSubmitInfo)
		info.GET("/weather", r.handleGetWeather)
		info.GET("/news", r.handleGetNews)
		info.GET("/ai-updates", r.handleGetAIUpdates)
	}

	// 状态管理相关路由
	status := router.Group("/status")
	{
		status.POST("/update", r.handleStatusUpdate)
		status.GET("/current", r.handleGetCurrentStatus)
		status.GET("/history", r.handleGetStatusHistory)
		status.GET("/metrics", r.handleGetSystemMetrics)
	}

	// 命令系统相关路由
	command := router.Group("/command")
	{
		command.POST("/execute", r.handleCommandExecute)
		command.GET("/history", r.handleGetCommandHistory)
		command.GET("/suggestions", r.handleGetCommandSuggestions)
	}

	// 配置管理相关路由
	config := router.Group("/config")
	{
		config.GET("", r.handleGetConfig)
		config.PUT("", r.handleUpdateConfig)
		config.GET("/defaults", r.handleGetDefaultConfig)
	}

	// WebSocket连接
	router.GET("/ws", r.handleWebSocket)
}

// 处理动画动作
func (r *Router) handleAvatarAction(c *gin.Context) {
	var req domain.AvatarActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "无效的请求参数",
				Details: err.Error(),
			},
		})
		return
	}

	response, err := r.service.ExecuteAvatarAction(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrActionNotAllowed,
				Message: "执行动作失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    response,
	})
}

// 获取可用动作
func (r *Router) handleGetAvailableActions(c *gin.Context) {
	actions := r.service.GetAvailableActions(c.Request.Context())
	
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"actions": actions,
			"emotions": []domain.Emotion{
				domain.EmotionHappy,
				domain.EmotionCurious,
				domain.EmotionSleepy,
				domain.EmotionExcited,
				domain.EmotionCalm,
				domain.EmotionAlert,
			},
		},
	})
}

// 获取信息收集
func (r *Router) handleGetInfoCollections(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5")
	categories := c.Query("categories")
	
	collections, err := r.service.GetInfoCollections(c.Request.Context(), limit, categories)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrInfoNotFound,
				Message: "获取信息失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    collections,
	})
}

// 更新状态
func (r *Router) handleStatusUpdate(c *gin.Context) {
	var req domain.StatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "无效的请求参数",
				Details: err.Error(),
			},
		})
		return
	}

	response, err := r.service.UpdateStatus(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrServerError,
				Message: "更新状态失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    response,
	})
}

// 执行命令
func (r *Router) handleCommandExecute(c *gin.Context) {
	var req domain.CommandExecuteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "无效的请求参数",
				Details: err.Error(),
			},
		})
		return
	}

	response, err := r.service.ExecuteCommand(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrCommandFailed,
				Message: "执行命令失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    response,
	})
}

// 获取配置
func (r *Router) handleGetConfig(c *gin.Context) {
	config, err := r.service.GetConfig(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrConfigInvalid,
				Message: "获取配置失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    config,
	})
}

// 更新配置
func (r *Router) handleUpdateConfig(c *gin.Context) {
	var config domain.OpenClawConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "无效的配置数据",
				Details: err.Error(),
			},
		})
		return
	}

	err := r.service.UpdateConfig(c.Request.Context(), config)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error": domain.ErrorResponse{
				Code:    domain.ErrConfigInvalid,
				Message: "更新配置失败",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "配置更新成功",
	})
}

// WebSocket处理
func (r *Router) handleWebSocket(c *gin.Context) {
	// 这里应该实现WebSocket升级和处理逻辑
	// 由于篇幅限制，这里只提供框架
	
	c.JSON(501, gin.H{
		"success": false,
		"error": domain.ErrorResponse{
			Code:    "NOT_IMPLEMENTED",
			Message: "WebSocket功能尚未实现",
		},
	})
}

// 其他处理函数（简化实现）
func (r *Router) handleGetAvatarStatus(c *gin.Context) {
	// 实现获取头像状态
}

func (r *Router) handleSubmitInfo(c *gin.Context) {
	// 实现提交信息
}

func (r *Router) handleGetWeather(c *gin.Context) {
	// 实现获取天气
}

func (r *Router) handleGetNews(c *gin.Context) {
	// 实现获取新闻
}

func (r *Router) handleGetAIUpdates(c *gin.Context) {
	// 实现获取AI更新
}

func (r *Router) handleGetCurrentStatus(c *gin.Context) {
	// 实现获取当前状态
}

func (r *Router) handleGetStatusHistory(c *gin.Context) {
	// 实现获取状态历史
}

func (r *Router) handleGetSystemMetrics(c *gin.Context) {
	// 实现获取系统指标
}

func (r *Router) handleGetCommandHistory(c *gin.Context) {
	// 实现获取命令历史
}

func (r *Router) handleGetCommandSuggestions(c *gin.Context) {
	// 实现获取命令建议
}

func (r *Router) handleGetDefaultConfig(c *gin.Context) {
	// 实现获取默认配置
}