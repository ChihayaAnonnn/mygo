package domain

import "time"

// ==================== KnowledgeNode ====================

// Node 知识图谱中的顶层节点，是所有知识内容的根抽象
type Node struct {
	ID             string     // 知识节点的全局唯一标识（UUID）
	NodeType       NodeType   // 节点类型（blog / note / paper / concept 等）
	Title          string     // 面向人类的知识标题
	Summary        string     // 短摘要，可由 AI 自动生成或人工维护
	Status         NodeStatus // 知识状态（draft / published / archived）
	Confidence     *float32   // 对该知识结论的主观置信度（0–1）
	CurrentVersion int        // 当前生效的 Markdown 版本号
	CreatedAt      time.Time  // 节点创建时间
	UpdatedAt      time.Time  // 节点最近一次元信息更新
}

// ==================== KnowledgeVersion ====================

// Version Markdown 内容的权威快照（Source of Truth）
type Version struct {
	ID        string    // 版本记录唯一标识（UUID）
	NodeID    string    // 所属知识节点 ID
	Version   int       // 版本号，从 1 开始单调递增
	ContentMd string    // 完整 Markdown 内容（权威文本）
	CreatedAt time.Time // 该版本生成时间
}

// ==================== KnowledgeChunk ====================

// Chunk AI 处理的最小单元（Embedding / RAG / Reasoning）
type Chunk struct {
	ID          string    // Chunk 唯一标识（UUID）
	NodeID      string    // 所属知识节点
	Version     int       // 来源的 Markdown 版本号
	HeadingPath string    // 该 Chunk 所属的标题层级路径（如 H1/H2）
	Content     string    // Chunk 的纯文本内容
	TokenCount  *int      // 该 Chunk 的 Token 数（用于模型预算）
	ChunkIndex  *int      // 在同一版本中的顺序编号
	CreatedAt   time.Time // Chunk 创建时间
}

// ==================== KnowledgeEdge ====================

// Edge 知识图谱中的边（Edge）
type Edge struct {
	ID        string    // 边的唯一标识（UUID）
	FromNode  string    // 起始知识节点
	ToNode    string    // 指向的知识节点
	EdgeType  EdgeType  // 关系类型
	CreatedAt time.Time // 关系创建时间
}

// ==================== KnowledgeEmbedding ====================

// Embedding Chunk 的语义向量表示
type Embedding struct {
	ChunkID   string    // 对应的 Chunk ID（一对一）
	Embedding []float32 // 语义向量（维度与模型绑定，如 1536）
	Model     string    // 生成该向量的模型名称
	CreatedAt time.Time // 向量生成时间
}
