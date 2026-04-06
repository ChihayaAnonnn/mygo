import { Link } from "react-router-dom";
import { motion } from "framer-motion";

const works = [
  {
    title: "OpenClaw Showcase",
    status: "Live",
    summary: "OpenClaw 数字伙伴的实时展示与交互入口，包含状态、动作和信息面板。",
    href: "/openclaw",
    cta: "Open showcase",
  },
  {
    slug: "signal-garden",
    title: "Signal Garden",
    status: "Planned",
    summary: "一个偏视觉叙事的实验入口，未来会放入动态场景、角色化信息与多媒体片段。",
    cta: "Open entry",
  },
  {
    slug: "companion-protocol",
    title: "Companion Protocol",
    status: "Prototype",
    summary: "一个为数字角色、陪伴感界面和 AI 交互预留的作品入口。",
    cta: "Open entry",
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
