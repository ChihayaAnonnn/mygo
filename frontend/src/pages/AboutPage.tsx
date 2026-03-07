export default function AboutPage() {
  return (
    <div>
      <div className="page-header">
        <span className="eyebrow">About</span>
        <h1>关于我</h1>
      </div>

      <div className="prose-card">
        <div className="markdown-body">
          <p>
            你好，我是一名关注工程实践与产品体验的开发者。这个网站主要记录我的学习过程、项目复盘和个人思考。
          </p>
          <p>
            目前博客处于第一版阶段，重点是保证内容更新稳定、阅读体验清晰，再逐步迭代功能。
          </p>
          <hr />
          <p>
            欢迎通过邮件与我交流：
            <a href="mailto:hello@example.com">hello@example.com</a>
          </p>
        </div>
      </div>
    </div>
  );
}
