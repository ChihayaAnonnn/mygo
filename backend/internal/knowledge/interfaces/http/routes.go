package http

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册知识相关路由
func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	knowledge := r.Group("/knowledge")
	{
		// 知识 CRUD
		knowledge.POST("", h.CreateKnowledge)
		knowledge.GET("", h.ListKnowledge)
		knowledge.GET("/:id", h.GetKnowledge)
		knowledge.PUT("/:id", h.UpdateKnowledgeMeta)
		knowledge.POST("/:id/archive", h.ArchiveKnowledge)

		// 版本管理
		knowledge.POST("/:id/versions", h.CreateVersion)
		knowledge.GET("/:id/versions", h.ListVersions)
		knowledge.GET("/:id/versions/latest", h.GetLatestVersion)

		// 应用层操作
		knowledge.POST("/:id/publish", h.PublishKnowledge)
		knowledge.POST("/:id/rebuild-index", h.RebuildIndex)

		// 语义搜索
		knowledge.POST("/search", h.Search)
	}
}
