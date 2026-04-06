import { motion } from "framer-motion";

const heroLines = [
  "I code. I design. I adapt.",
  "I create digital",
  "experiences that",
  "make sense.",
];

export default function Hero() {
  return (
    <section id="intro" className="hero-section" aria-labelledby="intro-title">
      <motion.div
        className="hero-copy"
        initial="hidden"
        animate="visible"
        variants={{
          hidden: {},
          visible: { transition: { staggerChildren: 0.18, delayChildren: 0.2 } },
        }}
      >
        <motion.p
          className="section-kicker"
          variants={{
            hidden: { opacity: 0, y: 18 },
            visible: { opacity: 1, y: 0, transition: { duration: 0.5 } },
          }}
        >
          Intro
        </motion.p>

        <h1 id="intro-title" className="hero-title" aria-label={heroLines.join(" ")}>
          {heroLines.map((line, lineIndex) => (
            <span key={line} className="hero-line">
              {line.split("").map((char, charIndex) => (
                <motion.span
                  key={`${lineIndex}-${charIndex}-${char}`}
                  className="hero-char"
                  variants={{
                    hidden: { opacity: 0, y: 30 },
                    visible: {
                      opacity: 1,
                      y: 0,
                      transition: {
                        duration: 0.36,
                        delay: lineIndex * 0.2 + charIndex * 0.028,
                      },
                    },
                  }}
                >
                  {char === " " ? "\u00A0" : char}
                </motion.span>
              ))}
            </span>
          ))}
        </h1>
      </motion.div>
    </section>
  );
}
