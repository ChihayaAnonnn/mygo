import { Link, NavLink, Outlet } from "react-router-dom";
import { posts } from "../content/posts";

export default function RootLayout() {
  return (
    <div className="app-layout">
      <aside className="app-sidebar">
        <Link to="/" className="sidebar-brand">
          my blog
        </Link>

        <p className="sidebar-section-label">导航</p>
        <nav className="sidebar-nav">
          <NavLink to="/" end>
            首页
          </NavLink>
          <NavLink to="/posts">文章</NavLink>
          <NavLink to="/about">关于</NavLink>
        </nav>

        <hr className="sidebar-divider" />

        <p className="sidebar-section-label">统计</p>
        <div style={{ fontSize: "0.82rem", color: "var(--text-muted)", lineHeight: "1.9" }}>
          <div>
            <span style={{ color: "var(--accent)", fontWeight: 600 }}>{posts.length}</span> 篇文章
          </div>
        </div>

        <div className="sidebar-footer">© {new Date().getFullYear()} My Blog</div>
      </aside>

      <main className="app-main">
        <Outlet />
      </main>
    </div>
  );
}
