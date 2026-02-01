# User 领域模块

用户认证与会话管理。

## 目录结构

```
user/
├── domain/
│   ├── model.go        # User 实体
│   ├── repository.go   # UserRepository, SessionCache 接口
│   ├── service.go      # UserService 接口
│   └── types.go        # 错误定义
│
├── application/
│   └── app_service.go  # 注册、登录、登出实现
│
├── infra/
│   ├── persistence/
│   │   ├── user_po.go
│   │   └── user_repo.go
│   └── cache/
│       └── session_cache.go
│
└── interfaces/http/
    ├── handler.go
    ├── routes.go
    └── dto.go
```

## API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/users/register | 用户注册 |
| POST | /api/users/login | 用户登录 |
| POST | /api/users/logout | 用户登出 |
| GET | /api/users/:id | 获取用户 |

## 领域模型

```go
type User struct {
    ID       int64
    UserID   int64
    Username string
    Email    string
    Password string
    Avatar   string
}
```

## 服务接口

```go
type UserService interface {
    Register(ctx, username, email, password) (*User, error)
    Login(ctx, username, password) (sessionID, *User, error)
    Logout(ctx, sessionID) error
    GetUserByID(ctx, id) (*User, error)
}
```
