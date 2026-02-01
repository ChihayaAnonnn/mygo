# 知识库数据库 Schema 设计文档（Knowledge Node 方案）

## 1. 设计背景与目标

本系统的目标并非传统博客或 CMS，而是一个 **可逐步演进为 AI‑Native 的个人知识系统（Personal Knowledge Base）**。因此，数据库设计从一开始就需要满足以下要求：

* **Knowledge-first，而非 Blog-first**
* 支持 Markdown 作为人类可编辑的知识源
* 支持 AI 消费（Chunk / Embedding / RAG）
* 支持知识之间的显式关系（Knowledge Graph）
* 在单机 PostgreSQL 环境下即可稳定运行，并可平滑演进

在命名与建模上，系统采用 **Knowledge Node（知识节点）** 作为顶层抽象，避免 `entity / post / document` 等易产生歧义的概念。

---

## 2. 核心设计原则

### 2.1 命名原则

* 避免与 DDD、ORM、AI 中已有高频术语冲突
* 数据表命名直接反映知识系统语义
* 为未来知识图谱（Graph）与 Agent 推理预留空间

### 2.2 权责划分

* **PostgreSQL**：权威事实来源（Schema / Markdown 快照 / Embedding）
* **文件系统**：编辑友好型源文件、恢复与工具支持
* **Redis（可选）**：缓存、异步任务加速

---

## 3. 数据模型总览

| 层级        | 表名                             | 作用               |
| --------- | ------------------------------ | ---------------- |
| 知识顶层      | `knowledge_nodes`              | 知识系统中的顶层节点       |
| 内容版本      | `knowledge_versions`           | Markdown 的权威版本快照 |
| AI 单元     | `knowledge_chunks`             | AI 处理的最小单元       |
| 语义向量      | `knowledge_embeddings`         | Chunk 对应的向量表示    |
| 知识关系      | `knowledge_edges`              | 知识节点之间的关系        |
| 辅助分类      | `tags` / `knowledge_node_tags` | 弱语义标签            |
| AI 任务（可选） | `ai_tasks`                     | AI 处理任务追踪        |

---

## 4. 表结构设计说明

本节在概念说明的基础上，**给出对应的 PostgreSQL DDL**，作为该知识系统的权威数据库定义。

---

### 4.1 knowledge_nodes —— 知识节点（核心表）

**语义定位**：

> 知识图谱中的一个节点，是所有知识内容的根抽象。

**字段说明**：

| 字段              | 类型          | 说明                                             |
| --------------- | ----------- | ---------------------------------------------- |
| id              | UUID        | 知识节点的全局唯一标识，是所有关联的锚点                           |
| node_type       | VARCHAR(32) | 节点类型，用于区分知识形态（blog / note / paper / concept 等） |
| title           | TEXT        | 面向人类的知识标题，允许自由修改                               |
| summary         | TEXT        | 短摘要，可由 AI 自动生成或人工维护                            |
| status          | VARCHAR(16) | 知识状态（draft / published / archived）             |
| confidence      | REAL        | 对该知识结论的主观置信度（0–1），服务于研究型知识                     |
| current_version | INT         | 当前生效的 Markdown 版本号                             |
| created_at      | TIMESTAMP   | 节点创建时间                                         |
| updated_at      | TIMESTAMP   | 节点最近一次元信息更新                                    |

**典型 node_type**：

* `blog`
* `note`
* `paper`
* `concept`
* `experiment`
* `code`

**DDL**：

```sql
CREATE TABLE knowledge_nodes (
    id UUID PRIMARY KEY,
    node_type VARCHAR(32) NOT NULL,
    title TEXT NOT NULL,
    summary TEXT,
    status VARCHAR(16) DEFAULT 'draft',
    confidence REAL,
    current_version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_knowledge_nodes_type ON knowledge_nodes(node_type);
CREATE INDEX idx_knowledge_nodes_status ON knowledge_nodes(status);
```

---

### 4.2 knowledge_versions —— Markdown 版本表

**语义定位**：

> Markdown 内容的权威快照（Source of Truth）。

**字段说明**：

| 字段         | 类型        | 说明                   |
| ---------- | --------- | -------------------- |
| id         | UUID      | 版本记录唯一标识             |
| node_id    | UUID      | 所属知识节点 ID            |
| version    | INT       | 版本号，从 1 开始单调递增       |
| content_md | TEXT      | 完整 Markdown 内容（权威文本） |
| created_at | TIMESTAMP | 该版本生成时间              |

**DDL**：

```sql
CREATE TABLE knowledge_versions (
    id UUID PRIMARY KEY,
    node_id UUID NOT NULL REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    version INT NOT NULL,
    content_md TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(node_id, version)
);

CREATE INDEX idx_knowledge_versions_node ON knowledge_versions(node_id);
```

---

### 4.3 knowledge_chunks —— 知识分块表

**语义定位**：

> AI 处理的最小单元（Embedding / RAG / Reasoning）。

**字段说明**：

| 字段           | 类型        | 说明                         |
| ------------ | --------- | -------------------------- |
| id           | UUID      | Chunk 唯一标识                 |
| node_id      | UUID      | 所属知识节点                     |
| version      | INT       | 来源的 Markdown 版本号           |
| heading_path | TEXT      | 该 Chunk 所属的标题层级路径（如 H1/H2） |
| content      | TEXT      | Chunk 的纯文本内容               |
| token_count  | INT       | 该 Chunk 的 Token 数（用于模型预算）  |
| chunk_index  | INT       | 在同一版本中的顺序编号                |
| created_at   | TIMESTAMP | Chunk 创建时间                 |

**DDL**：

```sql
CREATE TABLE knowledge_chunks (
    id UUID PRIMARY KEY,
    node_id UUID NOT NULL REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    version INT NOT NULL,
    heading_path TEXT,
    content TEXT NOT NULL,
    token_count INT,
    chunk_index INT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_knowledge_chunks_node ON knowledge_chunks(node_id);
CREATE INDEX idx_knowledge_chunks_version ON knowledge_chunks(node_id, version);
```

---

### 4.4 knowledge_embeddings —— 向量表

**语义定位**：

> Chunk 的语义向量表示。

**字段说明**：

| 字段         | 类型          | 说明                |
| ---------- | ----------- | ----------------- |
| chunk_id   | UUID        | 对应的 Chunk ID（一对一） |
| embedding  | VECTOR      | 语义向量（维度与模型绑定）     |
| model      | VARCHAR(64) | 生成该向量的模型名称        |
| created_at | TIMESTAMP   | 向量生成时间            |

**DDL**：

```sql
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE knowledge_embeddings (
    chunk_id UUID PRIMARY KEY REFERENCES knowledge_chunks(id) ON DELETE CASCADE,
    embedding VECTOR(1536) NOT NULL,
    model VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_knowledge_embeddings_vector
ON knowledge_embeddings
USING ivfflat (embedding vector_cosine_ops);
```

---

### 4.5 knowledge_edges —— 知识关系表

**语义定位**：

> 知识图谱中的边（Edge）。

**字段说明**：

| 字段         | 类型          | 说明                                         |
| ---------- | ----------- | ------------------------------------------ |
| id         | UUID        | 边的唯一标识                                     |
| from_node  | UUID        | 起始知识节点                                     |
| to_node    | UUID        | 指向的知识节点                                    |
| edge_type  | VARCHAR(64) | 关系类型（cites / derives_from / contradicts 等） |
| created_at | TIMESTAMP   | 关系创建时间                                     |

**DDL**：

```sql
CREATE TABLE knowledge_edges (
    id UUID PRIMARY KEY,
    from_node UUID NOT NULL REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    to_node UUID NOT NULL REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    edge_type VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_knowledge_edges_from ON knowledge_edges(from_node);
CREATE INDEX idx_knowledge_edges_to ON knowledge_edges(to_node);
CREATE INDEX idx_knowledge_edges_type ON knowledge_edges(edge_type);
```

---

### 4.6 Tags（弱语义分类）

**语义定位**：

> 人工或 AI 辅助的弱语义分类体系。

**字段说明**：

**tags 表**：

| 字段   | 类型          | 说明       |
| ---- | ----------- | -------- |
| id   | UUID        | 标签唯一标识   |
| name | VARCHAR(64) | 标签名称（唯一） |

**knowledge_node_tags 表**：

| 字段      | 类型   | 说明      |
| ------- | ---- | ------- |
| node_id | UUID | 知识节点 ID |
| tag_id  | UUID | 标签 ID   |

**DDL**：

```sql
CREATE TABLE tags (
    id UUID PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE knowledge_node_tags (
    node_id UUID REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (node_id, tag_id)
);
```

---

### 4.7 ai_tasks（可选）

**语义定位**：

> AI 处理流程的可追溯任务记录。

**字段说明**：

| 字段         | 类型          | 说明                                      |
| ---------- | ----------- | --------------------------------------- |
| id         | UUID        | AI 任务唯一标识                               |
| node_id    | UUID        | 关联的知识节点                                 |
| version    | INT         | 任务对应的 Markdown 版本                       |
| task_type  | VARCHAR(32) | 任务类型（chunk / embedding / summary 等）     |
| status     | VARCHAR(16) | 任务状态（pending / running / done / failed） |
| created_at | TIMESTAMP   | 任务创建时间                                  |
| updated_at | TIMESTAMP   | 最近一次状态更新                                |

**DDL**：

```sql
CREATE TABLE ai_tasks (
    id UUID PRIMARY KEY,
    node_id UUID REFERENCES knowledge_nodes(id) ON DELETE CASCADE,
    version INT,
    task_type VARCHAR(32),
    status VARCHAR(16),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

---

## 5. Markdown 与 AI 的生命周期

1. 用户创建 / 更新知识节点
2. 写入 `knowledge_versions`
3. 同步写入文件系统（current.md / vN.md）
4. 触发 AI 处理（异步）：

   * Chunk 切分
   * Embedding 生成
   * Summary / Edge 建议

**核心原则**：

> 同步路径最短，AI 全异步。

---

## 6. 演进路径保证

该 Schema 支持以下演进而无需重构：

* 从博客 → 知识库
* 从全文检索 → 向量检索（RAG）
* 从弱关联 → 知识图谱
* 从单机 → 服务化 AI Agent

---

## 7. 总结

本设计不是为“写文章”服务，而是为：

> **构建一个长期可演进、可计算、可被 AI 理解与利用的个人知识系统。**

Knowledge Node 方案在命名、语义与扩展性上，为未来 2–3 年的 AI 能力升级预留了稳定底座。
