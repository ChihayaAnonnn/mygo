package domain

import "time"

// ==================== KnowledgeNode ====================

// Node 知识图谱中的顶层节点，是所有知识内容的根抽象
type Node struct {
	ID             int64      // 数据库自增主键（内部使用）
	NodeID         string     // 知识节点的全局唯一标识（UUID，对外暴露）
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
	ID        int64     // 数据库自增主键（内部使用）
	VersionID string    // 版本记录唯一标识（UUID，对外暴露）
	NodeID    string    // 所属知识节点 UUID
	Version   int       // 版本号，从 1 开始单调递增
	ContentMd string    // 完整 Markdown 内容（权威文本）
	CreatedAt time.Time // 该版本生成时间
}

// ==================== KnowledgeChunk ====================

// Chunk AI 处理的最小单元（Embedding / RAG / Reasoning）
type Chunk struct {
	ID          int64     // 数据库自增主键（内部使用）
	ChunkID     string    // Chunk 唯一标识（UUID，对外暴露）
	NodeID      string    // 所属知识节点 UUID
	Version     int       // 来源的 Markdown 版本号
	HeadingPath string    // 该 Chunk 所属的标题层级路径（如 H1/H2）
	Content     string    // Chunk 的纯文本内容
	TokenCount  *int      // 该 Chunk 的 Token 数（用于模型预算）
	ChunkIndex  *int      // 在同一版本中的顺序编号
	CreatedAt   time.Time // Chunk 创建时间
}

// ==================== KnowledgeEdge（增强版）====================

// Edge 文档级关系边（Knowledge Node 之间），支持双时态
type Edge struct {
	ID             int64      // 数据库自增主键（内部使用）
	EdgeID         string     // 边的唯一标识（UUID，对外暴露）
	FromNode       string     // 起始知识节点 UUID
	ToNode         string     // 指向的知识节点 UUID
	EdgeType       EdgeType   // 关系类型
	Name           string     // 人类可读的关系描述
	Weight         float32    // 关系强度（0-1）
	Source         EdgeSource // 关系来源（manual / llm_extracted / rule_based）
	EpisodeID      string     // 来源 Episode UUID（可选）
	ValidFrom      *time.Time // 事实生效时间（Event Time）
	ValidUntil     *time.Time // 事实失效时间（NULL = 当前有效）
	InvalidatedAt  *time.Time // 被后续信息推翻的时间
	Metadata       map[string]any // 灵活扩展字段
	CreatedAt      time.Time  // 关系创建时间（Ingestion Time）
}

// ==================== KnowledgeEmbedding ====================

// Embedding Chunk 的语义向量表示
type Embedding struct {
	ID        int64     // 数据库自增主键（内部使用）
	ChunkID   string    // 对应的 Chunk UUID（一对一，唯一约束）
	Embedding []float32 // 语义向量（维度与模型绑定，如 1536）
	Model     string    // 生成该向量的模型名称
	CreatedAt time.Time // 向量生成时间
}

// ==================== Episode（数据摄入事件）====================

// Episode 数据摄入事件，记录每一次数据进入系统的上下文（溯源基础）
type Episode struct {
	ID                int64            // 数据库自增主键（内部使用）
	EpisodeID         string           // Episode 唯一标识（UUID，对外暴露）
	Name              string           // Episode 名称
	SourceType        EpisodeSourceType // 来源类型（text / message / json / markdown）
	SourceDescription string           // 来源描述
	Content           string           // 原始内容
	SourceNodeID      string           // 关联的 Knowledge Node UUID（可选）
	ReferenceTime     time.Time        // 事件实际发生时间（Event Time）
	CreatedAt         time.Time        // 摄入系统时间（Ingestion Time）
}

// ==================== Entity（实体节点）====================

// Entity 从文档/对话中抽取的细粒度实体（知识图谱的原子节点）
type Entity struct {
	ID            int64      // 数据库自增主键（内部使用）
	EntityID      string     // 实体唯一标识（UUID，对外暴露）
	EntityType    EntityType // 实体类型（person / organization / concept 等）
	Name          string     // 实体名称
	Summary       string     // 实体描述/摘要
	NameEmbedding []float32  // 实体名称的语义向量（用于语义搜索去重）
	GroupID       string     // 所属 Community UUID（可选）
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 最近更新时间
}

// ==================== EntityEdge（实体关系）====================

// EntityEdge 实体间的语义关系，以三元组形式存储，支持双时态
type EntityEdge struct {
	ID             int64      // 数据库自增主键（内部使用）
	EdgeID         string     // 边的唯一标识（UUID，对外暴露）
	FromEntity     string     // 起始实体 UUID（Subject）
	ToEntity       string     // 目标实体 UUID（Object）
	EdgeType       string     // 关系类型（Predicate 分类标签）
	Name           string     // 人类可读的关系名称（如 "founded"）
	Fact           string     // 完整事实描述（如 "Pradip founded FutureSmart AI in 2020"）
	FactEmbedding  []float32  // 事实描述的语义向量
	Weight         float32    // 关系强度（0-1）
	EpisodeID      string     // 来源 Episode UUID
	ValidFrom      *time.Time // 事实生效时间（Event Time）
	ValidUntil     *time.Time // 事实失效时间（NULL = 当前有效）
	InvalidatedAt  *time.Time // 被后续信息推翻的时间
	CreatedAt      time.Time  // 边创建时间（Ingestion Time）
}

// ==================== Community（社区聚类）====================

// Community 实体的层次化社区分组，由社区检测算法（如 Leiden）生成
type Community struct {
	ID              int64     // 数据库自增主键（内部使用）
	CommunityID     string    // 社区唯一标识（UUID，对外暴露）
	Name            string    // 社区名称
	Summary         string    // 社区摘要（由 LLM 生成）
	Level           int       // 层级（0 = 最细粒度）
	ParentCommunity string    // 上级社区 UUID（层次化结构，可选）
	CreatedAt       time.Time // 创建时间
	UpdatedAt       time.Time // 最近更新时间
}
