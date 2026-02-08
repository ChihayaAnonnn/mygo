# 文档命名规范

## 命名规则

### 1. 根目录全局性文档（全大写）

位于项目根目录或子项目根目录的重要配置和说明文档，使用全大写命名：

```text
/workspace/mygo/
├── AGENT.md              ✓ 项目级 AI 配置
├── ERRORS.md             ✓ 全局错误记录
├── README.md             ✓ 项目说明
└── CONTRIBUTING.md       ✓ 贡献指南（如有）

frontend/
├── DEVELOPMENT.md        ✓ 前端开发指南
└── DESIGN_SYSTEM.md      ✓ 前端设计规范
```

### 2. 技术文档（小写+下划线）

位于 `docs/` 目录下的技术细节、架构设计、接口文档，使用小写+下划线：

```text
backend/docs/
├── architecture.md                  ✓ 架构概览
├── development.md                   ✓ 开发指南
├── writing_guide.md                 ✓ 文档编写规范
├── naming_convention.md             ✓ 命名规范
├── knowledge_schema_design.md       ✓ 数据库设计
└── knowledge_interface_design.md    ✓ 接口设计
```

### 3. 模块级文档（全大写）

位于模块/包根目录的说明文档，使用全大写：

```text
backend/internal/knowledge/
└── README.md             ✓ 模块说明

backend/internal/user/
└── README.md             ✓ 模块说明
```

## 规则总结

| 位置 | 命名规则 | 示例 | 理由 |
| --- | --- | --- | --- |
| 项目根目录 | 全大写 | `AGENT.md` | 醒目，表示重要性 |
| 子项目根目录 | 全大写 | `frontend/DEVELOPMENT.md` | 相当于子项目的"门面" |
| docs/ 目录 | 小写+下划线 | `architecture.md` | 技术文档数量多，统一小写易读 |
| 模块根目录 | 全大写 | `knowledge/README.md` | 模块入口说明 |

## 重命名历史

2026-01-31: 统一文档命名规范

- `Agent.md` → `AGENT.md`
- `backend/docs/Architecture.md` → `architecture.md`
- `backend/docs/DEVELOPMENT.md` → `development.md`
- `backend/docs/WRITING_GUIDE.md` → `writing_guide.md`
