# 文档编写规范

## 格式要求

### Markdown 格式

所有项目文档必须使用 **Markdown** 格式编写（`.md` 文件）。

### 图表绘制

使用 **Mermaid** 语法绘制图表，支持的图表类型包括：

- 流程图 (Flowchart)
- 时序图 (Sequence Diagram)
- 类图 (Class Diagram)
- 状态图 (State Diagram)
- ER 图 (Entity Relationship Diagram)
- 甘特图 (Gantt Chart)

#### 示例：流程图

```mermaid
flowchart TD
    A[开始] --> B{是否需要创建?}
    B -->|是| C[创建实体]
    B -->|否| D[更新实体]
    C --> E[保存到数据库]
    D --> E
    E --> F[返回结果]
```

#### 示例：时序图

```mermaid
sequenceDiagram
    participant Client
    participant Handler
    participant Service
    participant Repository
    
    Client->>Handler: HTTP Request
    Handler->>Service: CreateKnowledge()
    Service->>Repository: Save()
    Repository-->>Service: Entity
    Service-->>Handler: Result
    Handler-->>Client: HTTP Response
```

#### 示例：ER 图

```mermaid
erDiagram
    KNOWLEDGE_NODE ||--o{ KNOWLEDGE_VERSION : has
    KNOWLEDGE_NODE ||--o{ KNOWLEDGE_CHUNK : contains
    KNOWLEDGE_VERSION ||--o{ KNOWLEDGE_CHUNK : generates
    
    KNOWLEDGE_NODE {
        uuid id PK
        string title
        string status
        timestamp created_at
    }
    
    KNOWLEDGE_VERSION {
        uuid id PK
        uuid node_id FK
        text content
        int version_number
    }
```

## 编写原则

- **简洁明了**：避免冗长和重复
- **结构清晰**：使用标题层级组织内容
- **代码示例**：关键概念提供代码示例
- **图文并茂**：复杂流程使用 Mermaid 图表说明
