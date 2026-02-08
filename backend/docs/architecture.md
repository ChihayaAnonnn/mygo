# 后端架构文档

> Data Plane Service for AI Agent。采用 **Clean Architecture / DDD-lite** 架构，按领域模块组织代码。

## 目录结构

```bash
/mygo/backend
├── cmd/                        # 入口程序
│   ├── server/main.go          # HTTP 服务入口
│   └── migrate/main.go         # 数据迁移入口
│
├── internal/
│   ├── bootstrap/              # 启动引导（app/http/worker/migrate）
│   ├── config/                 # 配置管理
│   ├── infra/                  # 共享基础设施（DB/Redis）
│   ├── server/                 # 全局路由聚合
│   ├── user/                   # ★ User 领域模块
│   └── knowledge/              # ★ Knowledge 领域模块（数据 CRUD + 向量检索）
│
├── deployments/                # 部署配置
└── docs/                       # 文档
```

## Clean Architecture 领域模块结构

每个领域模块采用以下四层结构：

```text
internal/<domain>/
├── domain/                 # 领域层（最稳定，零外部依赖）
│   ├── model.go            # 领域模型/实体
│   ├── repository.go       # Repository 接口
│   ├── service.go          # Service 接口
│   └── types.go            # 错误、枚举、Command/Query
│
├── application/            # 用例层（编排业务流程）
│   └── app_service.go
│
├── infra[structure]/       # 基础设施层（技术实现）
│   ├── persistence/        # 数据库（PO + Repo 实现）
│   ├── cache/              # 缓存
│   └── ...                 # 其他外部服务
│
└── interfaces/             # 接口适配层（协议转换）
    └── http/               # handler/routes/dto
```

## 依赖规则

```text
interfaces → domain ← application
                ↑
              infra
```

- **依赖方向始终指向 domain 层**（依赖倒置）
- `domain` 定义接口，`infra` 提供实现
- `application` 编排业务，`interfaces` 处理协议

## 领域模块

| 模块 | 说明 | 文档 |
|------|------|------|
| `user/` | 用户认证与会话管理 | [README](../internal/user/README.md) |
| `knowledge/` | 数据 CRUD（Node/Version/Chunk/Embedding）、向量检索 | [README](../internal/knowledge/README.md) |

## 环境变量

| 变量名     | 默认值 | 描述 |
|-----------|--------|------|
| PORT      | 8080   | 服务端口 |
| GIN_MODE  | debug  | Gin 运行模式 |
| PG_DSN    | postgres://... | PostgreSQL DSN |
| REDIS_URL | redis://... | Redis URL |

## 新增领域模块

1. 创建 `internal/<domain>/` 四层目录结构
2. 在 `bootstrap/app.go` 添加模块初始化
3. 在 `server/router.go` 注册路由
4. 创建 `internal/<domain>/README.md` 文档
