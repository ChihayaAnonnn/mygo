import { Link } from "react-router-dom";

export default function NotFoundPage() {
  return (
    <main className="page">
      <div className="bg" aria-hidden="true">
        <div className="glow glow-1" />
        <div className="glow glow-2" />
      </div>

      <section className="card">
        <header className="header">
          <p className="kicker">mygo</p>
          <span className="pill" aria-label="状态">
            404
          </span>
        </header>

        <h1 className="title">
          页面不存在
          <span className="titleSub">你访问的地址可能已变更或尚未上线</span>
        </h1>

        <p className="lead">点击下方按钮返回首页。</p>

        <footer className="footer">
          <div className="hint">
            <span className="dot" aria-hidden="true" />
            <span>返回后你会看到当前的占位页。</span>
          </div>
          <Link className="link" to="/">
            回到首页 →
          </Link>
        </footer>
      </section>
    </main>
  );
}
