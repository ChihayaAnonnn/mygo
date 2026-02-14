package domain

import (
	"errors"
	"time"
)

// ==================== 领域 ID 类型（业务 UUID） ====================
// 这些 ID 类型代表对外暴露的业务标识（UUID），用于 API 和跨系统引用
// 数据库内部使用自增 int64 作为主键，提升索引性能

// KnowledgeID 知识节点的业务唯一标识（UUID）
type KnowledgeID string

// VersionID 版本记录的业务唯一标识（UUID）
type VersionID string

// ChunkID 分块的业务唯一标识（UUID）
type ChunkID string

// EpisodeID 数据摄入事件的业务唯一标识（UUID）
type EpisodeID string

// EntityID 实体的业务唯一标识（UUID）
type EntityID string

// EntityEdgeID 实体关系的业务唯一标识（UUID）
type EntityEdgeID string

// CommunityID 社区的业务唯一标识（UUID）
type CommunityID string

// FilePath 文件路径
type FilePath string

// EmbeddingVector 向量表示
type EmbeddingVector []float32

// ==================== 领域错误 ====================

var (
	ErrKnowledgeNotFound  = errors.New("knowledge not found")
	ErrVersionNotFound    = errors.New("knowledge version not found")
	ErrChunkNotFound      = errors.New("knowledge chunk not found")
	ErrEdgeNotFound       = errors.New("knowledge edge not found")
	ErrEmbeddingNotFound  = errors.New("knowledge embedding not found")
	ErrEpisodeNotFound    = errors.New("episode not found")
	ErrEntityNotFound     = errors.New("entity not found")
	ErrEntityEdgeNotFound = errors.New("entity edge not found")
	ErrCommunityNotFound  = errors.New("community not found")
	ErrInvalidInput       = errors.New("invalid input")
)

// ==================== 节点类型 ====================

// NodeType 定义知识节点类型
type NodeType string

const (
	NodeTypeBlog       NodeType = "blog"
	NodeTypeNote       NodeType = "note"
	NodeTypePaper      NodeType = "paper"
	NodeTypeConcept    NodeType = "concept"
	NodeTypeExperiment NodeType = "experiment"
	NodeTypeCode       NodeType = "code"
)

// ==================== 节点状态 ====================

// NodeStatus 定义知识节点状态
type NodeStatus string

const (
	NodeStatusDraft     NodeStatus = "draft"
	NodeStatusPublished NodeStatus = "published"
	NodeStatusArchived  NodeStatus = "archived"
)

// ==================== 边类型 ====================

// EdgeType 定义文档级关系类型
type EdgeType string

const (
	EdgeTypeCites       EdgeType = "cites"        // 引用
	EdgeTypeDerivesFrom EdgeType = "derives_from" // 派生自
	EdgeTypeContradicts EdgeType = "contradicts"  // 矛盾
	EdgeTypeRelatesTo   EdgeType = "relates_to"   // 相关
	EdgeTypeExtends     EdgeType = "extends"      // 扩展
	EdgeTypeSupports    EdgeType = "supports"     // 支持
)

// ==================== 边来源 ====================

// EdgeSource 定义关系的来源
type EdgeSource string

const (
	EdgeSourceManual       EdgeSource = "manual"        // 人工创建
	EdgeSourceLLMExtracted EdgeSource = "llm_extracted" // LLM 自动抽取
	EdgeSourceRuleBased    EdgeSource = "rule_based"    // 规则推导
)

// ==================== Episode 来源类型 ====================

// EpisodeSourceType 定义 Episode 的数据来源类型
type EpisodeSourceType string

const (
	EpisodeSourceText     EpisodeSourceType = "text"     // 非结构化文本
	EpisodeSourceMessage  EpisodeSourceType = "message"  // 对话消息
	EpisodeSourceJSON     EpisodeSourceType = "json"     // 结构化 JSON
	EpisodeSourceMarkdown EpisodeSourceType = "markdown" // Markdown 文档
)

// ==================== 实体类型 ====================

// EntityType 定义实体类型（可由 Agent 自定义扩展，此处为常见预设）
type EntityType string

const (
	EntityTypePerson       EntityType = "person"       // 人物
	EntityTypeOrganization EntityType = "organization" // 组织
	EntityTypeConcept      EntityType = "concept"      // 概念
	EntityTypeTechnology   EntityType = "technology"   // 技术
	EntityTypeEvent        EntityType = "event"        // 事件
	EntityTypeLocation     EntityType = "location"     // 地点
	EntityTypeProduct      EntityType = "product"      // 产品
)

// ==================== Command 对象（文档层） ====================

// CreateKnowledgeCmd 创建知识的命令
type CreateKnowledgeCmd struct {
	NodeType NodeType // 节点类型
	Title    string   // 标题
	Summary  string   // 摘要（可选）
}

// UpdateKnowledgeMetaCmd 更新知识元信息的命令
type UpdateKnowledgeMetaCmd struct {
	ID         KnowledgeID // 知识 ID
	Title      *string     // 标题（可选更新）
	Summary    *string     // 摘要（可选更新）
	Status     *NodeStatus // 状态（可选更新）
	Confidence *float32    // 置信度（可选更新）
}

// CreateVersionCmd 创建版本的命令
type CreateVersionCmd struct {
	KnowledgeID KnowledgeID // 所属知识 ID
	ContentMd   string      // Markdown 内容
}

// ==================== Command 对象（图谱层） ====================

// CreateEpisodeCmd 创建 Episode 的命令
type CreateEpisodeCmd struct {
	Name              string            // Episode 名称
	SourceType        EpisodeSourceType // 来源类型
	SourceDescription string            // 来源描述（可选）
	Content           string            // 原始内容
	SourceNodeID      KnowledgeID       // 关联的 Knowledge Node（可选）
	ReferenceTime     time.Time         // 事件实际发生时间
}

// CreateEntityCmd 创建实体的命令
type CreateEntityCmd struct {
	EntityType    EntityType      // 实体类型
	Name          string          // 实体名称
	Summary       string          // 实体描述（可选）
	NameEmbedding EmbeddingVector // 名称的语义向量（可选）
	GroupID       CommunityID     // 所属社区（可选）
}

// UpdateEntityCmd 更新实体的命令
type UpdateEntityCmd struct {
	ID            EntityID        // 实体 ID
	Name          *string         // 名称（可选更新）
	Summary       *string         // 描述（可选更新）
	NameEmbedding EmbeddingVector // 语义向量（可选更新）
	GroupID       *CommunityID    // 社区（可选更新）
}

// CreateEntityEdgeCmd 创建实体关系的命令
type CreateEntityEdgeCmd struct {
	FromEntity    EntityID        // 起始实体
	ToEntity      EntityID        // 目标实体
	EdgeType      string          // 关系类型
	Name          string          // 关系名称
	Fact          string          // 完整事实描述
	FactEmbedding EmbeddingVector // 事实的语义向量（可选）
	Weight        float32         // 关系强度（0-1）
	EpisodeID     EpisodeID       // 来源 Episode（可选）
	ValidFrom     *time.Time      // 事实生效时间（可选）
}

// CreateCommunityCmd 创建社区的命令
type CreateCommunityCmd struct {
	Name            string      // 社区名称
	Summary         string      // 社区摘要
	Level           int         // 层级
	ParentCommunity CommunityID // 上级社区（可选）
	MemberIDs       []EntityID  // 成员实体 ID 列表
}

// ==================== Query 对象 ====================

// KnowledgeQuery 知识查询条件
type KnowledgeQuery struct {
	NodeType *NodeType   // 按类型筛选
	Status   *NodeStatus // 按状态筛选
	Keyword  *string     // 关键词搜索（标题/摘要）
	Offset   int         // 分页偏移
	Limit    int         // 分页大小
}

// EpisodeQuery Episode 查询条件
type EpisodeQuery struct {
	SourceType   *EpisodeSourceType // 按来源类型筛选
	SourceNodeID *KnowledgeID       // 按关联 Node 筛选
	Offset       int                // 分页偏移
	Limit        int                // 分页大小
}

// EntityQuery 实体查询条件
type EntityQuery struct {
	EntityType *EntityType  // 按实体类型筛选
	GroupID    *CommunityID // 按社区筛选
	Keyword    *string      // 关键词搜索（名称/摘要）
	Offset     int          // 分页偏移
	Limit      int          // 分页大小
}

// CommunityQuery 社区查询条件
type CommunityQuery struct {
	Level  *int // 按层级筛选
	Offset int  // 分页偏移
	Limit  int  // 分页大小
}
