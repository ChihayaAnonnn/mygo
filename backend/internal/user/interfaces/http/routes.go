package http

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册用户相关路由
func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	users := r.Group("/users")
	{
		users.POST("/register", h.Register)
		users.POST("/login", h.Login)
		users.POST("/logout", h.Logout)
		users.GET("/:id", h.GetUser)
	}
}
