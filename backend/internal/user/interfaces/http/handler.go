package http

import (
	"errors"
	"net/http"
	"strconv"

	"mygo/internal/user/domain"

	"github.com/gin-gonic/gin"
)

// Handler 用户 HTTP 处理器
type Handler struct {
	userService domain.UserService
}

// NewHandler 构造函数
func NewHandler(userService domain.UserService) *Handler {
	return &Handler{userService: userService}
}

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 响应辅助函数
func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func fail(c *gin.Context, httpCode int, code int, message string) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
	})
}

// Register 用户注册
// POST /api/users/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	user, err := h.userService.Register(c.Request.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserAlreadyExists):
			fail(c, http.StatusConflict, 409, "user already exists")
		case errors.Is(err, domain.ErrInvalidInput):
			fail(c, http.StatusBadRequest, 400, "invalid input")
		default:
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	resp := &RegisterResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
	}
	success(c, resp)
}

// Login 用户登录
// POST /api/users/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	sessionID, user, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidCredentials):
			fail(c, http.StatusUnauthorized, 401, "invalid username or password")
		case errors.Is(err, domain.ErrInvalidInput):
			fail(c, http.StatusBadRequest, 400, "invalid input")
		default:
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	resp := &LoginResponse{
		SessionID: sessionID,
		UserID:    user.UserID,
		Username:  user.Username,
	}
	success(c, resp)
}

// Logout 用户登出
// POST /api/users/logout
func (h *Handler) Logout(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		fail(c, http.StatusBadRequest, 400, "session id required")
		return
	}

	if err := h.userService.Logout(c.Request.Context(), sessionID); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, nil)
}

// GetUser 获取用户信息
// GET /api/users/:id
func (h *Handler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid user id")
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			fail(c, http.StatusNotFound, 404, "user not found")
		} else {
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	resp := &UserResponse{
		ID:       user.ID,
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	success(c, resp)
}
