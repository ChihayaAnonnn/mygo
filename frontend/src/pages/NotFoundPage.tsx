import { Link } from "react-router-dom";

export default function NotFoundPage() {
  return (
    <section className="site-container page-section">
      <h1>404</h1>
      <p>The page does not exist. The URL may be incorrect.</p>
      <Link className="text-link" to="/">
        返回首页
      </Link>
    </section>
  );
}
