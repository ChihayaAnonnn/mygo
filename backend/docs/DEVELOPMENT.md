# 后端开发指南

## 快速开始

### 运行服务

在 `backend/` 目录下执行：

```bash
go run cmd/server/main.go
```

### 运行测试

```bash
go test ./...
```

### 依赖管理

整理依赖：

```bash
go mod tidy
```

## 数据库迁移

使用 GORM AutoMigrate 自动同步数据模型到数据库。

### 执行迁移

```bash
go run cmd/migrate/main.go
```

### 预览模式（Dry-run）

在事务中执行迁移并回滚，仅预览 SQL 不实际改动数据库：

```bash
go run cmd/migrate/main.go --dry-run
```

### 添加新数据模型

1. 在 `internal/models/po/` 目录下创建新的 PO 文件，例如 `post_po.go`
2. 编辑 `cmd/migrate/main.go`，在 `models` 切片中注册新模型：

```go
var models = []any{
    &po.UserPO{},
    &po.PostPO{},  // 新增
}
```

3. 执行迁移：`go run cmd/migrate/main.go`

## 部署与环境

- **Docker Compose**: 配置文件位于 `deployments/compose.yaml`。
