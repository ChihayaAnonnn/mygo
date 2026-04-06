import { motion } from "framer-motion";
import { ArrowUpRight } from "lucide-react";

const picks = [
  {
    title: "Are.na",
    url: "https://www.are.na/",
    kind: "Website",
    source: "are.na",
    note:
      "Saved for the calm structure and the feeling that browsing itself can be a creative act.",
    tags: ["archive", "visual culture", "curation"],
    revisitHint: "I keep returning to it for the pace.",
  },
  {
    title: "A List Apart",
    url: "https://alistapart.com/",
    kind: "Article",
    source: "alistapart.com",
    note:
      "Still useful when I want writing about the web to feel opinionated, practical, and readable at the same time.",
    tags: ["writing", "web", "editorial"],
    revisitHint: "Often revisit it for structure and tone.",
  },
  {
    title: "Linear",
    url: "https://linear.app/",
    kind: "Product",
    source: "linear.app",
    note:
      "A reminder that motion, hierarchy, and density can feel precise without becoming cold.",
    tags: ["product", "motion", "ui systems"],
    revisitHint: "Saved for its interaction discipline.",
  },
  {
    title: "The Creative Independent",
    url: "https://thecreativeindependent.com/",
    kind: "Essay",
    source: "thecreativeindependent.com",
    note:
      "I like the warmth of the interviews and how creative process gets treated as something lived, not just performed.",
    tags: ["interview", "process", "creative life"],
    revisitHint: "Returned for the voice and honesty.",
  },
];

const containerVariants = {
  hidden: {},
  visible: {
    transition: {
      staggerChildren: 0.12,
    },
  },
};

const itemVariants = {
  hidden: { opacity: 0, y: 30 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.55,
      ease: [0.22, 1, 0.36, 1] as const,
    },
  },
};

export default function MyPicks() {
  return (
    <motion.section
      id="picks"
      className="content-section picks-section"
      aria-labelledby="picks-title"
      initial={{ opacity: 0, y: 44 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true, amount: 0.2 }}
      transition={{ duration: 0.7 }}
    >
      <div className="section-headline picks-headline">
        <p className="section-kicker">My Picks</p>
        <h2 id="picks-title">一些收集到的好东西.</h2>
      </div>

      <motion.div
        className="picks-list"
        variants={containerVariants}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.2 }}
      >
        {picks.map((pick) => (
          <motion.a
            key={pick.title}
            href={pick.url}
            className="pick-entry"
            target="_blank"
            rel="noreferrer"
            variants={itemVariants}
            whileHover={{ y: -4 }}
          >
            <div className="pick-meta">
              <span>{pick.kind}</span>
              <span>{pick.source}</span>
            </div>

            <div className="pick-body">
              <div className="pick-copy">
                <h3>{pick.title}</h3>
                <p>{pick.note}</p>
              </div>

              <div className="pick-footer">
                <div className="pick-tags" aria-label={`${pick.title} tags`}>
                  {pick.tags.map((tag) => (
                    <span key={tag}>{tag}</span>
                  ))}
                </div>
                <p className="pick-hint">{pick.revisitHint}</p>
              </div>
            </div>

            <span className="pick-open">
              Open
              <ArrowUpRight size={16} strokeWidth={2} />
            </span>
          </motion.a>
        ))}
      </motion.div>

      <div className="picks-outro">
        <span className="text-link">Browse all picks</span>
        <p>More links, notes, and internet keepsakes will gather here over time.</p>
      </div>
    </motion.section>
  );
}
