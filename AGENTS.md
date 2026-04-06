# 项目上下文配置 (Agent Context)

## 项目概述

这是一个属于我自己的 **个人数字游乐场**。
它是一个开放的个人网站与自由实验场，用来承载我想创造的一切：可以是文章、交互作品、视觉实验、AI 体验、多模态内容、数字人、二次元/动漫风格页面，或任何我感兴趣的表达形式。项目不预设唯一产品形态，重点是围绕我的兴趣、审美、表达欲与创造冲动持续演化。

### 核心设计共识

- **Creator-first**: 一切设计优先服务于我的兴趣、表达和创造欲，而不是迎合固定的网站范式。
- **Format-agnostic**: 它可以是文章、页面、角色、装置、交互叙事、AI 体验或混合媒介作品。
- **Imagination-friendly**: 参与开发的 Agent 应主动发挥想象力，与我讨论大胆但合适的想法，不要默认把项目收窄成普通博客或后台系统。
- **Evolvable**: 网站可以持续变化，不要求所有页面统一成单一结构；允许不同主题、不同媒介形式并存，只要整体体验是有意识设计过的。
- **前后端分离**: 前端负责体验、视觉、交互和内容呈现；后端保留独立服务与基础设施能力，为未来扩展提供空间。

## 技术架构与关键目录

### 数据存储 (Data Storage)

- **当前路径**: `frontend/posts/`
- **当前用途**: 这是现阶段的一种内容来源，用于存放 Markdown 文章。
- **设计态度**: 内容载体不受限于文章。未来也可以扩展为交互页面、媒体资源、角色设定、结构化数据、生成内容、实验性素材或其他任意形式的创作资产。

### 后端服务 (`backend/`)

- **技术栈**: Go, Gorm (ORM), Redis, Gin (HTTP)
- **架构**: Clean Architecture / DDD-lite，按领域模块组织代码
- **关键目录**:
  - `cmd/server/`: HTTP 服务入口
  - `cmd/migrate/`: 数据库迁移工具
  - `internal/bootstrap/`: 启动引导（app/http/worker/migrate）
  - `internal/server/`: 全局路由聚合
  - `internal/user/`: User 领域模块
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
  - `src/pages/`: 页面与独立体验入口
  - `src/content/`: 当前内容源与内容解析逻辑
  - `src/layouts/`: 站点级布局组件
  - `src/styles.css`: 全局视觉语言与页面风格基础

## 设计决策背景

- **前后端分离**: 允许前端和后端独立开发、部署和扩展
- **创作驱动**: 项目首先是表达与实验的容器，不要求围绕单一内容模型组织一切页面
- **形式开放**: 页面可以是阅读型、叙事型、交互型、角色型、视觉型或混合体验
- **渐进演化**: 允许先从一个具体形态开始，再逐步长成完全不同的东西，不因当前实现限制未来方向
- **Clean Architecture / DDD-lite**:
  - **领域驱动**: 按业务领域组织代码
  - **依赖倒置**: 领域层定义接口，基础设施层实现
  - **可测试性**: 通过接口隔离，可轻松 mock 依赖
  - **可替换性**: 更换技术实现不影响业务逻辑

## 文档索引

### 根目录

| 文档 | 说明 |
| --- | --- |
| [ERRORS.md](ERRORS.md) | 常见错误记录（**开发时注意查阅**） |

### 后端文档 (`backend/docs/`)

| 文档 | 说明 |
| --- | --- |
| [architecture.md](backend/docs/architecture.md) | 架构概览、目录结构、依赖规则 |
| [development.md](backend/docs/development.md) | 开发指南、快速开始 |
| [writing_guide.md](backend/docs/writing_guide.md) | 文档编写规范（**编写文档前必读**） |
| [naming_convention.md](backend/docs/naming_convention.md) | 文档命名规范 |

### 领域模块文档 (`backend/internal/`)

> 每个领域模块的设计文档放在其自身目录下（`backend/internal/<domain>/README.md`），而非集中在 `backend/docs/`。

| 文档 | 说明 |
| --- | --- |
| [user/README.md](backend/internal/user/README.md) | User 模块：API、模型、接口 |

### 前端文档 (`frontend/`)

| 文档 | 说明 |
| --- | --- |
| [DEVELOPMENT.md](frontend/DEVELOPMENT.md) | 前端开发指南 |
| [DESIGN_SYSTEM.md](frontend/DESIGN_SYSTEM.md) | UI 设计规范（**开发前必读**） |

## 开发环境

- **Docker Compose**:
  - `backend/deployments/compose.yaml`
  - `frontend/compose.yaml`

## 生产环境

- **访问地址**: `http://mygo.chat`
- **服务状态**: 已部署并运行正常
- **反向代理**: 使用 Nginx 进行域名访问

## Agent 行为准则

- **查阅常见错误**：开发前先查看 [ERRORS.md](ERRORS.md)，避免重复犯错
- **禁止扮演 Linter 角色**：不要手动修正代码格式或风格问题
- **依赖自动化工具**：专注于逻辑实现，格式问题留给工具链处理
- **遵循架构**：新增代码必须遵循 Clean Architecture 分层原则
- **文档规范**：编写文档时遵循 [writing_guide.md](backend/docs/writing_guide.md)，使用 Markdown 格式，图表使用 Mermaid
- **禁止使用 Emoji**：在文档、提交信息和代码注释中，除非用户明确要求，否则不要使用 Emoji 表情
