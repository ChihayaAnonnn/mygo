package domain

import "errors"

// ==================== 领域 ID 类型（业务 UUID） ====================
// 这些 ID 类型代表对外暴露的业务标识（UUID），用于 API 和跨系统引用
// 数据库内部使用自增 int64 作为主键，提升索引性能

// KnowledgeID 知识节点的业务唯一标识（UUID）
type KnowledgeID string

// VersionID 版本记录的业务唯一标识（UUID）
type VersionID string

// ChunkID 分块的业务唯一标识（UUID）
type ChunkID string

// FilePath 文件路径
type FilePath string

// EmbeddingVector 向量表示
type EmbeddingVector []float32

// ==================== 领域错误 ====================

var (
	ErrKnowledgeNotFound = errors.New("knowledge not found")
	ErrVersionNotFound   = errors.New("knowledge version not found")
	ErrChunkNotFound     = errors.New("knowledge chunk not found")
	ErrEdgeNotFound      = errors.New("knowledge edge not found")
	ErrEmbeddingNotFound = errors.New("knowledge embedding not found")
	ErrInvalidInput      = errors.New("invalid input")
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

// EdgeType 定义知识关系类型
type EdgeType string

const (
	EdgeTypeCites       EdgeType = "cites"        // 引用
	EdgeTypeDerivesFrom EdgeType = "derives_from" // 派生自
	EdgeTypeContradicts EdgeType = "contradicts"  // 矛盾
	EdgeTypeRelatesTo   EdgeType = "relates_to"   // 相关
	EdgeTypeExtends     EdgeType = "extends"      // 扩展
	EdgeTypeSupports    EdgeType = "supports"     // 支持
)

// ==================== Command 对象 ====================

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

// ==================== Query 对象 ====================

// KnowledgeQuery 知识查询条件
type KnowledgeQuery struct {
	NodeType *NodeType   // 按类型筛选
	Status   *NodeStatus // 按状态筛选
	Keyword  *string     // 关键词搜索（标题/摘要）
	Offset   int         // 分页偏移
	Limit    int         // 分页大小
}
