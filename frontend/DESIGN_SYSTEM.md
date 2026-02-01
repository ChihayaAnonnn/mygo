# MyGo 设计系统

> Agent 开发 UI 时的设计与色彩指南。所有页面和组件应遵循此规范。

## 设计理念

**碳灰色与高饱和中明度颜色的组合 —— 高端、经典、先锋摩登**

### 核心特点

- **稳定与跃动**: 碳灰提供稳定、中性的基底；高饱和中明度色彩是精心校准的感官焦点
- **清晰层级**: 碳灰区域让眼睛休息；高饱和色抓取视线，指向重要内容
- **现代与经典**: 碳灰自带工业风、高级灰经典感；高饱和色注入活力与当代气息

### 视觉质感

- **冷峻科技感**: 碳灰如磨砂金属，搭配青绿色产生精密仪器、未来界面的质感
- **可靠与创新**: 碳灰传递安全稳固；亮色体现智能高效

---

## 色彩系统

### 配色比例 (7:3 黄金比例)

| 角色 | 色值 | 占比 | 用途 |
|------|------|------|------|
| 主色 | `#052228` | 70% | 背景、大面积区域 |
| 副色 | `#9FE7E6` | 15% | 强调、按钮、焦点元素 |
| 辅助色 | `#156B8C` | 10% | 次要交互、hover、边框 |
| 点缀色 | `#E8FBFF` | 5% | 高亮文字、微光效果 |

### CSS 变量

```css
:root {
  /* 主色调 - 深碳灰/墨绿 (70%) */
  --bg-primary: #052228;
  --bg-secondary: #0a3038;
  --bg-elevated: #0d3a42;

  /* 副色调 - 青绿 (15%) */
  --accent: #9FE7E6;
  --accent-hover: #b5edec;
  --accent-muted: rgba(159, 231, 230, 0.6);

  /* 辅助色 - 深青蓝 (10%) */
  --secondary: #156B8C;
  --secondary-hover: #1a7fa6;

  /* 点缀色 - 极浅青白 (5%) */
  --highlight: #E8FBFF;
  --highlight-muted: rgba(232, 251, 255, 0.8);

  /* 文字色 */
  --text-primary: #E8FBFF;
  --text-secondary: rgba(232, 251, 255, 0.7);
  --text-muted: rgba(232, 251, 255, 0.5);

  /* 边框与阴影 */
  --border: rgba(159, 231, 230, 0.15);
  --border-strong: rgba(159, 231, 230, 0.3);
  --shadow: 0 24px 48px rgba(0, 0, 0, 0.4);
  --shadow-glow: 0 0 20px rgba(159, 231, 230, 0.15);
}
```

### 使用原则

1. **一个焦点**: 只用 `--accent` 作为主亮点，`--secondary` 降低面积使用
2. **始终使用 CSS 变量**，禁止硬编码颜色

---

## 字体系统

| 用途 | 字体 | 备选 |
|------|------|------|
| 标题 | `Fraunces` | `ui-serif, Georgia, serif` |
| 正文 | `IBM Plex Sans` | `ui-sans-serif, system-ui, sans-serif` |

```css
@import url("https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,600;9..144,700&family=IBM+Plex+Sans:wght@400;500;600&display=swap");
```

| 元素 | 字号 | 行高 | 字间距 |
|------|------|------|--------|
| 主标题 | `clamp(30px, 4.2vw, 52px)` | 1.06 | -0.02em |
| 副标题 | `clamp(14px, 1.6vw, 18px)` | 1.4 | 0.01em |
| 正文 | 16px | 1.7 | - |
| 小标签 | 12px | - | 0.08em |

---

## 视觉特效

### 深色玻璃态

```css
background: rgba(5, 34, 40, 0.85);
backdrop-filter: blur(16px);
border: 1px solid var(--border);
box-shadow: var(--shadow);
```

### 发光效果

```css
/* 元素发光 */
box-shadow: 0 0 20px rgba(159, 231, 230, 0.2), 0 0 40px rgba(159, 231, 230, 0.1);

/* 文字发光 */
text-shadow: 0 0 10px rgba(159, 231, 230, 0.5);
```

### 渐变

```css
/* 背景渐变 */
background: linear-gradient(180deg, #052228, #0a3038);

/* 强调渐变 */
background: linear-gradient(135deg, #9FE7E6, #156B8C);
```

---

## 组件规范

### 卡片

- 背景: `rgba(5, 34, 40, 0.85)` + `blur(16px)`
- 圆角: `16px`，内边距: `24px`
- 边框: `var(--border)`

### 按钮

```css
/* 主按钮 */
.btn-primary {
  background: var(--accent);
  color: var(--bg-primary);
  border-radius: 8px;
  padding: 12px 24px;
  box-shadow: var(--shadow-glow);
}
.btn-primary:hover {
  background: var(--accent-hover);
  box-shadow: 0 0 30px rgba(159, 231, 230, 0.3);
}

/* 次要按钮 */
.btn-secondary {
  background: transparent;
  color: var(--accent);
  border: 1px solid var(--border-strong);
}
```

### 标签/徽章

- 背景: `rgba(159, 231, 230, 0.15)`
- 颜色: `var(--accent)`
- 圆角: `999px`

### 链接

- 颜色: `var(--accent)`
- Hover: 添加下划线 + 微光 `text-shadow`

### 状态点

- 8px 圆点，颜色 `var(--accent)`，带发光 `box-shadow`

---

## 页面结构

```html
<main class="page">
  <div class="bg" aria-hidden="true">...</div>
  <section class="card">...</section>
</main>
```

```css
.page {
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 28px 18px;
  background: var(--bg-primary);
  color: var(--text-primary);
}
```

---

## 开发检查清单

- [ ] 使用 CSS 变量，不硬编码颜色
- [ ] 遵循 7:3 配色比例（碳灰 70% / 亮色 30%）
- [ ] 只用一种高饱和色作为主焦点
- [ ] 标题用 Fraunces，正文用 IBM Plex Sans
- [ ] 卡片应用深色玻璃态 + 发光效果
- [ ] 圆角: 卡片 16px / 按钮 8px / 徽章 999px
- [ ] 装饰元素加 `aria-hidden="true"`

---

## 快速参考

| 属性 | 值 |
|------|-----|
| 主背景 | `#052228` |
| 强调色 | `#9FE7E6` |
| 辅助色 | `#156B8C` |
| 高亮色 | `#E8FBFF` |
| 卡片圆角 | `16px` |
| 模糊度 | `blur(16px)` |
| 标题字体 | Fraunces |
| 正文字体 | IBM Plex Sans |
| 配色比例 | 70:15:10:5 |
