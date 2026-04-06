import { motion } from "framer-motion";

const contactItems = [
  { label: "Email", value: "hello@example.com", href: "mailto:hello@example.com" },
  { label: "GitHub", value: "@mygo", href: "https://github.com/" },
  { label: "X / Social", value: "@your-handle", href: "https://x.com/" },
];

export default function Contact() {
  return (
    <motion.section
      id="contact"
      className="content-section contact-section"
      aria-labelledby="contact-title"
      initial={{ opacity: 0, y: 44 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true, amount: 0.25 }}
      transition={{ duration: 0.7 }}
    >
      <div className="section-headline">
        <p className="section-kicker">Contact</p>
        <h2 id="contact-title">如果你对这些想法感兴趣，欢迎来和我一起继续把它们做出来。</h2>
      </div>

      <div className="contact-list">
        {contactItems.map((item) => (
          <motion.a
            key={item.label}
            href={item.href}
            className="contact-item"
            target={item.href.startsWith("http") ? "_blank" : undefined}
            rel={item.href.startsWith("http") ? "noreferrer" : undefined}
            initial={{ opacity: 0, y: 22 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true, amount: 0.4 }}
            transition={{ duration: 0.45 }}
          >
            <span className="contact-label">{item.label}</span>
            <span className="contact-value">{item.value}</span>
          </motion.a>
        ))}
      </div>
    </motion.section>
  );
}
