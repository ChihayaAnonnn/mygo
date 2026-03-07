import { Link } from "react-router-dom";
import { posts } from "../content/posts";

export default function PostsPage() {
  const [featured, second, ...rest] = posts;

  return (
    <div>
      <div className="page-header">
        <span className="eyebrow">Archive</span>
        <h1>全部文章</h1>
        <p>共 {posts.length} 篇，按发布时间倒序排列。</p>
      </div>

      {posts.length === 0 ? (
        <p style={{ color: "var(--text-muted)" }}>暂无文章。</p>
      ) : (
        <div className="magazine-grid">
          {featured && (
            <article className="post-card post-card-featured">
              <p className="post-meta">{featured.date}</p>
              <h2 className="post-card-title">
                <Link to={`/posts/${featured.slug}`}>{featured.title}</Link>
              </h2>
              <p className="post-card-summary">{featured.summary}</p>
              {featured.tags.length > 0 && (
                <div className="tag-list">
                  {featured.tags.map((tag) => (
                    <span key={tag} className="tag">
                      {tag}
                    </span>
                  ))}
                </div>
              )}
            </article>
          )}

          {second && (
            <article className="post-card">
              <p className="post-meta">{second.date}</p>
              <h2 className="post-card-title">
                <Link to={`/posts/${second.slug}`}>{second.title}</Link>
              </h2>
              <p className="post-card-summary">{second.summary}</p>
              {second.tags.length > 0 && (
                <div className="tag-list">
                  {second.tags.map((tag) => (
                    <span key={tag} className="tag">
                      {tag}
                    </span>
                  ))}
                </div>
              )}
            </article>
          )}

          {rest.map((post) => (
            <article key={post.slug} className="post-card">
              <p className="post-meta">{post.date}</p>
              <h2 className="post-card-title">
                <Link to={`/posts/${post.slug}`}>{post.title}</Link>
              </h2>
              <p className="post-card-summary">{post.summary}</p>
              {post.tags.length > 0 && (
                <div className="tag-list">
                  {post.tags.map((tag) => (
                    <span key={tag} className="tag">
                      {tag}
                    </span>
                  ))}
                </div>
              )}
            </article>
          ))}
        </div>
      )}
    </div>
  );
}
