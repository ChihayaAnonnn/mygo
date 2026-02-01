export default function ComingSoonPage() {
  return (
    <main className="page">
      <div className="bg" aria-hidden="true">
        <div className="glow glow-1" />
        <div className="glow glow-2" />
      </div>

      <section className="card">
        <header className="header">
          <p className="kicker">mygo</p>
          <span className="pill" aria-label="站点状态">
            开发中
          </span>
        </header>

        <h1 className="title">
          网站正在开发中
          <span className="titleSub">高端 · 经典 · 先锋摩登</span>
        </h1>

        <p className="lead">
          我在把它做成一个干净、克制、但细节丰富的个人网站。第一版上线前，这里先用作占位页。
        </p>

        <footer className="footer">
          <div className="hint">
            <span className="dot" aria-hidden="true" />
            <span>感谢你的到访，晚点再来。</span>
          </div>
          <a className="link" href="/" onClick={(e) => e.preventDefault()}>
            刷新看看 →
          </a>
        </footer>
      </section>
    </main>
  );
}
