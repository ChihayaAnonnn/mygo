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

## 部署与环境

- **Docker Compose**: 配置文件位于 `deployments/compose.yaml`。
