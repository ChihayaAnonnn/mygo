package domain

import "time"

// OpenClawStatus 表示OpenClaw的状态
type OpenClawStatus string

const (
	StatusOnline   OpenClawStatus = "online"
	StatusBusy     OpenClawStatus = "busy"
	StatusAway     OpenClawStatus = "away"
	StatusSleeping OpenClawStatus = "sleeping"
)

// AvatarAction 表示动画形象的动作
type AvatarAction string

const (
	ActionIdle   AvatarAction = "idle"
	ActionWave   AvatarAction = "wave"
	ActionDance  AvatarAction = "dance"
	ActionSleep  AvatarAction = "sleep"
	ActionEat    AvatarAction = "eat"
	ActionWork   AvatarAction = "work"
	ActionPlay   AvatarAction = "play"
	ActionReact  AvatarAction = "react"
	ActionSpeak  AvatarAction = "speak"
)

// Emotion 表示情绪状态
type Emotion string

const (
	EmotionHappy   Emotion = "happy"
	EmotionCurious Emotion = "curious"
	EmotionSleepy  Emotion = "sleepy"
	EmotionExcited Emotion = "excited"
	EmotionCalm    Emotion = "calm"
	EmotionAlert   Emotion = "alert"
)

// AvatarActionRequest 动画动作请求
type AvatarActionRequest struct {
	Action    AvatarAction           `json:"action"`
	Intensity float64                `json:"intensity"` // 0.0-1.0
	Duration  int                    `json:"duration"`  // 毫秒
	Emotion   Emotion                `json:"emotion"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// AvatarActionResponse 动画动作响应
type AvatarActionResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ActionID  string `json:"action_id"`
	Timestamp int64  `json:"timestamp"`
}

// WeatherInfo 天气信息
type WeatherInfo struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	Humidity    float64 `json:"humidity"`
	Forecast    string  `json:"forecast"`
	Icon        string  `json:"icon"`
	UpdatedAt   int64   `json:"updated_at"`
}

// NewsItem 新闻条目
type NewsItem struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	Summary     string    `json:"summary"`
	URL         string    `json:"url"`
	Category    string    `json:"category"`
	PublishedAt time.Time `json:"published_at"`
	Relevance   float64   `json:"relevance"` // 0.0-1.0
}

// AIUpdate AI发展动态
type AIUpdate struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Source      string    `json:"source"`
	Impact      string    `json:"impact"` // low/medium/high
	Date        time.Time `json:"date"`
}

// PersonalStats 个人统计数据
type PersonalStats struct {
	DailyTasksCompleted int       `json:"daily_tasks_completed"`
	MessagesProcessed   int       `json:"messages_processed"`
	SkillsUsed          []string  `json:"skills_used"`
	Mood                string    `json:"mood"`
	LastActive          time.Time `json:"last_active"`
}

// InfoCollections 信息收集
type InfoCollections struct {
	Weather   WeatherInfo   `json:"weather"`
	News      []NewsItem    `json:"news"`
	AIUpdates []AIUpdate    `json:"ai_updates"`
	Personal  PersonalStats `json:"personal"`
	Timestamp int64         `json:"timestamp"`
}

// StatusUpdateRequest 状态更新请求
type StatusUpdateRequest struct {
	Status   OpenClawStatus          `json:"status"`
	Activity string                  `json:"activity"`
	Mood     string                  `json:"mood"`
	Metrics  map[string]interface{}  `json:"metrics"`
	Message  string                  `json:"message"`
}

// StatusUpdateResponse 状态更新响应
type StatusUpdateResponse struct {
	Success      bool  `json:"success"`
	Acknowledged bool  `json:"acknowledged"`
	Timestamp    int64 `json:"timestamp"`
}

// CommandExecuteRequest 命令执行请求
type CommandExecuteRequest struct {
	Command    string                 `json:"command"`
	Parameters map[string]interface{} `json:"parameters"`
	Source     string                 `json:"source"` // web/tui/api
}

// CommandExecuteResponse 命令执行响应
type CommandExecuteResponse struct {
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Message     string      `json:"message"`
	ExecutionID string      `json:"execution_id"`
}

// WebSocket消息类型
type WebSocketMessageType string

const (
	MsgTypeStatusUpdate WebSocketMessageType = "status_update"
	MsgTypeAvatarAction WebSocketMessageType = "avatar_action"
	MsgTypeInfoUpdate   WebSocketMessageType = "info_update"
	MsgTypeNotification WebSocketMessageType = "notification"
	MsgTypeCommandResult WebSocketMessageType = "command_result"
)

// ClientMessage 客户端到服务器的消息
type ClientMessage struct {
	Type WebSocketMessageType `json:"type"`
	Data interface{}          `json:"data"`
}

// ServerMessage 服务器到客户端的消息
type ServerMessage struct {
	Type      WebSocketMessageType `json:"type"`
	Data      interface{}          `json:"data"`
	Timestamp int64                `json:"timestamp"`
}

// SubscriptionRequest 订阅请求
type SubscriptionRequest struct {
	Channel string `json:"channel"`
}

// Notification 系统通知
type Notification struct {
	Type    string `json:"type"`    // info/warning/error/success
	Message string `json:"message"`
	Title   string `json:"title,omitempty"`
}

// SystemMetrics 系统指标
type SystemMetrics struct {
	CPUUsage      float64 `json:"cpu_usage"`
	MemoryUsage   float64 `json:"memory_usage"`
	ActiveSessions int    `json:"active_sessions"`
	Uptime        int64   `json:"uptime"` // 秒
	ResponseTime  int     `json:"response_time"` // 毫秒
}

// OpenClawConfig 配置信息
type OpenClawConfig struct {
	Avatar struct {
		AutoActions         bool    `json:"auto_actions"`
		ReactionSensitivity float64 `json:"reaction_sensitivity"`
		DefaultEmotion      Emotion `json:"default_emotion"`
	} `json:"avatar"`
	
	InfoDisplay struct {
		RefreshInterval int      `json:"refresh_interval"` // 毫秒
		MaxNewsItems    int      `json:"max_news_items"`
		Categories      []string `json:"categories"`
	} `json:"info_display"`
	
	Notifications struct {
		Enabled bool `json:"enabled"`
		Sound   bool `json:"sound"`
		Desktop bool `json:"desktop"`
	} `json:"notifications"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// 错误代码
const (
	ErrInvalidAction     = "INVALID_ACTION"
	ErrActionNotAllowed  = "ACTION_NOT_ALLOWED"
	ErrInfoNotFound      = "INFO_NOT_FOUND"
	ErrCommandFailed     = "COMMAND_FAILED"
	ErrConfigInvalid     = "CONFIG_INVALID"
	ErrAuthRequired      = "AUTH_REQUIRED"
	ErrInvalidToken      = "INVALID_TOKEN"
	ErrRateLimited       = "RATE_LIMITED"
	ErrServerError       = "SERVER_ERROR"
)