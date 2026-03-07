import { Link, useParams } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { getPostBySlug } from "../content/posts";

export default function PostDetailPage() {
  const { slug } = useParams<{ slug: string }>();
  const post = slug ? getPostBySlug(slug) : undefined;

  if (!post) {
    return (
      <div>
        <div className="page-header">
          <h1>文章不存在</h1>
          <p>你访问的文章未找到，可能已被移除或链接错误。</p>
        </div>
        <Link to="/posts" className="btn-secondary">
          返回文章列表
        </Link>
      </div>
    );
  }

  return (
    <div>
      <div className="page-header">
        <span className="eyebrow">
          <Link to="/posts" className="text-link" style={{ fontSize: "inherit", fontWeight: "inherit", letterSpacing: "inherit", textTransform: "inherit" }}>
            ← 文章
          </Link>
        </span>
        <p className="post-meta" style={{ marginTop: "8px" }}>
          {post.date}
        </p>
        <h1>{post.title}</h1>
        <p style={{ color: "var(--text-secondary)", marginTop: "6px", marginBottom: "0" }}>
          {post.summary}
        </p>
        {post.tags.length > 0 && (
          <div className="tag-list">
            {post.tags.map((tag) => (
              <span key={tag} className="tag">
                {tag}
              </span>
            ))}
          </div>
        )}
      </div>

      <div className="prose-card">
        <div className="markdown-body">
          <ReactMarkdown remarkPlugins={[remarkGfm]}>{post.content}</ReactMarkdown>
        </div>
      </div>
    </div>
  );
}
