import { Link } from "react-router-dom";

export default function OpenClawRetiredPage() {
  return (
    <section className="openclaw-retired-page">
      <div className="openclaw-retired-shell content-section">
        <p className="section-kicker">OpenClaw</p>
        <h1 className="openclaw-retired-title">OpenClaw 交互界面</h1>
        <p className="openclaw-retired-lead">
          由于 OpenClaw 存在安全风险，所以 OpenClaw 不再提供访问，ByeBye OpenClaw。
        </p>

        <div className="openclaw-retired-console" aria-label="OpenClaw interface preview">
          <div className="openclaw-retired-console-head">
            <span>Connection</span>
            <span className="openclaw-retired-badge">Disabled</span>
          </div>
          <p className="openclaw-retired-note">
            所有交互能力已被禁用，当前仅保留静态说明信息。
          </p>
        </div>

        <Link to="/" className="inline-link">
          返回首页
        </Link>
      </div>
    </section>
  );
}
