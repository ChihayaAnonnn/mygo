import { Link } from "react-router-dom";
import { motion } from "framer-motion";

interface WorkItem {
  slug: string;
  title: string;
  status: string;
  summary: string;
  cta: string;
  href?: string;
}

const works: WorkItem[] = [
  {
    slug: "agent-chat",
    title: "Dual LLM Agent Chat",
    status: "Live",
    summary: "Dual LLM Agent Chat 交互界面，包含聊天、背景 Agent 日志与 Workspace 文件浏览。",
    cta: "Open chat",
    href: "/agent-chat",
  },
  {
    slug: "openclaw-archived",
    title: "OpenClaw Interactive",
    status: "Archived",
    summary: "保留交互界面预览。由于出现安全问题，该页面已停止对外提供实际访问能力。",
    cta: "View notice",
    href: "/openclaw",
  }
];

export default function Work() {
  return (
    <motion.section
      id="work"
      className="content-section work-section"
      aria-labelledby="work-title"
      initial={{ opacity: 0, y: 44 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true, amount: 0.2 }}
      transition={{ duration: 0.7 }}
    >
      <div className="section-headline">
        <p className="section-kicker">Work</p>
        <h2 id="work-title">一个试验田.</h2>
      </div>

      <div className="work-grid">
        {works.map((work) => (
          <motion.div
            key={work.href ?? work.slug}
            initial={{ opacity: 0, y: 24 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true, amount: 0.3 }}
            transition={{ duration: 0.5 }}
          >
            <Link to={work.href ?? `/works/${work.slug}`} className="work-card">
              <div className="work-card-meta">
                <span>{work.status}</span>
              </div>
              <h3>{work.title}</h3>
              <p>{work.summary}</p>
              <span className="work-card-link">{work.cta}</span>
            </Link>
          </motion.div>
        ))}
      </div>
    </motion.section>
  );
}
