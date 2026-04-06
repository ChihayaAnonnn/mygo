import { Link, useParams } from "react-router-dom";

export default function WorkDetailPage() {
  const { slug } = useParams<{ slug: string }>();

  return (
    <section className="work-detail-page">
      <div className="work-detail-inner">
        <p className="section-kicker">Work</p>
        <h1>{slug?.replace(/-/g, " ") || "coming soon"}</h1>
        <p>
          这个入口已经预留好，未来会承载新的创意页面、交互实验或叙事体验。
        </p>
        <Link to="/" className="inline-link">
          返回首页
        </Link>
      </div>
    </section>
  );
}
