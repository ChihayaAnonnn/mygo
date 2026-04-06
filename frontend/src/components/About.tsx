import { motion } from "framer-motion";

export default function About() {
  return (
    <motion.section
      id="about"
      className="content-section about-section"
      aria-labelledby="about-title"
      initial={{ opacity: 0, y: 44 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true, amount: 0.25 }}
      transition={{ duration: 0.7 }}
    >
      <div className="section-headline">
        <p className="section-kicker">About</p>
        <h2 id="about-title">我希望网站和我本人一样，保持变化、热情和好奇心。</h2>
      </div>

      <div className="about-copy">
        <p>
          我对代码、视觉、角色、叙事、动漫、AI 与各种数字表达形式都有兴趣。
          比起定义自己是什么，我更在意一件事：能不能把脑海里的想法变成一个真实可感的体验。
        </p>
        <p>
          所以这里不会被限制为单一用途的网站。它可以是作品集、实验室、角色空间、情绪档案，
          或者一个暂时还说不清名字的新东西。
        </p>
      </div>
    </motion.section>
  );
}
