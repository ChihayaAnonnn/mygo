# 前端开发指南

## 快速开始

### 安装依赖与启动

在 `frontend/` 目录下执行：

```bash
npm install
npm run dev
```

### 如何关闭（停止开发服务器）

- **前台运行（最常见）**：在运行 `npm run dev` 的终端窗口里按下 `Ctrl + C` 即可停止。
- **端口仍被占用（少见）**：如果停止后 `5173` 端口仍被占用，可手动结束进程（Linux）：

```bash
lsof -i :5173
kill -9 <PID>
```

### 构建生产版本

```bash
npm run build
```

## 部署与环境

- **Docker Compose**: 配置文件位于 `compose.yaml`。
