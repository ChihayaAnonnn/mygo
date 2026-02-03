# 项目上下文配置 (Agent Context)

## 项目概述

这是一个 **AI-Native 个人知识库系统**。
核心目标是构建以 **Knowledge** 为中心，支持 Markdown 编辑、多版本演进、AI 处理（Chunk / Embedding / RAG）的长期可演进系统。

### 核心设计共识

- **Database as Source of Truth**: 数据库是知识事实源，Markdown 完整内容存于 DB，文件系统中的 `.md` 仅是可重建的投影。
- **Knowledge Node**: 顶层抽象，支持版本管理、知识图谱（Graph）和 AI 推理。
- **三层模型**:
  - DB：事实层（版本、关系、Chunk、Embedding）
  - FS：投影层（人类可读 Markdown，存储于 `/workspace/data/knowledge`）
  - AI：推理层（RAG / Agent，全异步可重建）

## 技术架构与关键目录

### 数据存储 (Data Storage)

- **路径**: `/workspace/data/knowledge`
- **用途**: 存储 Markdown 文件及其他数据文件的持久化目录（作为 DB 的文件系统投影）。

### 后端服务 (`backend/`)

- **技术栈**: Go, Gorm (ORM), Redis, Gin (HTTP)
- **架构**: Clean Architecture / DDD-lite，按领域模块组织代码
- **关键目录**:
  - `cmd/server/`: HTTP 服务入口
  - `cmd/migrate/`: 数据库迁移工具
  - `internal/bootstrap/`: 启动引导（app/http/worker/migrate）
  - `internal/server/`: 全局路由聚合
  - `internal/user/`: User 领域模块
  - `internal/knowledge/`: Knowledge 领域模块
  - `internal/infra/`: 共享基础设施 (DB, Redis)

#### Clean Architecture 领域模块结构

每个领域模块采用以下分层结构：

```text
internal/<domain>/
├── domain/                 # 领域层（最稳定，零外部依赖）
│   ├── model.go            # 领域模型/实体
│   ├── repository.go       # Repository 接口定义
│   ├── service.go          # Service 接口定义
│   └── types.go            # 领域错误、值对象、枚举
│
├── application/            # 用例层（编排领域对象完成业务流程）
│   └── app_service.go      # 应用服务，实现 domain.Service 接口
│
├── infra[structure]/       # 基础设施层（技术实现，最不稳定）
│   ├── persistence/        # 数据库持久化（PO + Repo 实现）
│   └── cache/              # 缓存实现
│
└── interfaces/             # 接口适配层（协议转换）
    └── http/
        ├── handler.go
        ├── routes.go
        └── dto.go
```

#### 依赖规则（核心原则）

```text
interfaces → domain ← application
                ↑
              infra
```

- **依赖方向始终指向 domain 层**（依赖倒置）
- `domain` 定义接口，`infra` 提供实现
- `application` 编排业务，`interfaces` 处理协议

#### 新增领域模块步骤

1. 创建 `internal/<domain>/` 四层目录结构
2. 在 `bootstrap/app.go` 添加模块初始化
3. 在 `server/router.go` 注册路由
4. 创建 `internal/<domain>/README.md` 文档

### 前端应用 (`frontend/`)

- **技术栈**: React, TypeScript, Vite, Tailwind CSS
- **关键目录**:
  - `src/pages/`: 页面组件
  - `src/App.tsx`: 应用根组件

## 设计决策背景

- **前后端分离**: 允许前端和后端独立开发、部署和扩展
- **Clean Architecture / DDD-lite**:
  - **领域驱动**: 按业务领域组织代码
  - **依赖倒置**: 领域层定义接口，基础设施层实现
  - **可测试性**: 通过接口隔离，可轻松 mock 依赖
  - **可替换性**: 更换技术实现不影响业务逻辑

## 文档索引

### 根目录

| 文档 | 说明 |
|------|------|
| [ERRORS.md](ERRORS.md) | 常见错误记录（**开发时注意查阅**） |

### 后端文档 (`backend/docs/`)

| 文档 | 说明 |
|------|------|
| [Architecture.md](backend/docs/Architecture.md) | 架构概览、目录结构、依赖规则 |
| [DEVELOPMENT.md](backend/docs/DEVELOPMENT.md) | 开发指南、快速开始 |
| [knowledge_schema_design.md](backend/docs/knowledge_schema_design.md) | Knowledge 数据库设计 |
| [knowledge_interface_design.md](backend/docs/knowledge_interface_design.md) | Knowledge 接口设计 |

### 领域模块文档 (`backend/internal/`)

| 文档 | 说明 |
|------|------|
| [user/README.md](backend/internal/user/README.md) | User 模块：API、模型、接口 |
| [knowledge/README.md](backend/internal/knowledge/README.md) | Knowledge 模块：API、服务、接口 |

### 前端文档 (`frontend/`)

| 文档 | 说明 |
|------|------|
| [DEVELOPMENT.md](frontend/DEVELOPMENT.md) | 前端开发指南 |
| [DESIGN_SYSTEM.md](frontend/DESIGN_SYSTEM.md) | UI 设计规范（**开发前必读**） |

## 开发环境

- **Docker Compose**:
  - `backend/deployments/compose.yaml`
  - `frontend/compose.yaml`

## Agent 行为准则

- **查阅常见错误**：开发前先查看 [ERRORS.md](ERRORS.md)，避免重复犯错
- **禁止扮演 Linter 角色**：不要手动修正代码格式或风格问题
- **依赖自动化工具**：专注于逻辑实现，格式问题留给工具链处理
- **遵循架构**：新增代码必须遵循 Clean Architecture 分层原则
- **禁止使用 Emoji**：在文档、提交信息和代码注释中，除非用户明确要求，否则不要使用 Emoji 表情
