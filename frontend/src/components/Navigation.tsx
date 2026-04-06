import { useEffect, useState } from "react";
import { Menu, X } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import { ThemeToggle } from "./ThemeToggle";

const sections = [
  { id: "intro", label: "Intro" },
  { id: "work", label: "Work" },
  { id: "picks", label: "Picks" },
  { id: "about", label: "About" },
  { id: "contact", label: "Contact" },
];

export default function Navigation() {
  const [activeSection, setActiveSection] = useState("intro");
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  useEffect(() => {
    const elements = document.querySelectorAll<HTMLElement>("section[id]");

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            setActiveSection(entry.target.id);
          }
        });
      },
      { rootMargin: "-45% 0px -45% 0px", threshold: 0 }
    );

    elements.forEach((element) => observer.observe(element));

    return () => observer.disconnect();
  }, []);

  const handleJump = (sectionId: string) => {
    const section = document.getElementById(sectionId);
    section?.scrollIntoView({ behavior: "smooth", block: "start" });
    setIsMenuOpen(false);
  };

  return (
    <>
      <motion.nav
        className="site-nav-desktop"
        aria-label="Primary navigation"
        initial={{ opacity: 0, x: -18 }}
        animate={{ opacity: 1, x: 0 }}
        transition={{ duration: 0.7, delay: 0.2 }}
      >
        <a
          className="site-mark"
          href="#intro"
          onClick={(event) => {
            event.preventDefault();
            handleJump("intro");
          }}
        >
          mygo
        </a>

        <div className="site-nav-links">
          {sections.map((section, index) => (
            <motion.button
              key={section.id}
              type="button"
              className={`site-nav-link${activeSection === section.id ? " is-active" : ""}`}
              onClick={() => handleJump(section.id)}
              initial={{ opacity: 0, x: -14 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ duration: 0.45, delay: 0.24 + index * 0.05 }}
            >
              {section.label}
            </motion.button>
          ))}
        </div>

        <motion.div
          className="site-nav-actions"
          initial={{ opacity: 0, x: -14 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.45, delay: 0.5 }}
        >
          <ThemeToggle />
        </motion.div>
      </motion.nav>

      <motion.div
        className="site-nav-mobile"
        initial={{ opacity: 0, y: -16 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.65, delay: 0.2 }}
      >
        <a
          className="site-mark"
          href="#intro"
          onClick={(event) => {
            event.preventDefault();
            handleJump("intro");
          }}
        >
          mygo
        </a>

        <div className="site-nav-mobile-actions">
          <ThemeToggle />
          <button
            type="button"
            className="site-menu-button"
            onClick={() => setIsMenuOpen((value) => !value)}
            aria-label={isMenuOpen ? "Close navigation" : "Open navigation"}
          >
            {isMenuOpen ? <X size={18} /> : <Menu size={18} />}
          </button>
        </div>
      </motion.div>

      <AnimatePresence>
        {isMenuOpen && (
          <motion.div
            className="site-mobile-panel"
            initial={{ opacity: 0, y: -20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -20 }}
            transition={{ duration: 0.28 }}
          >
            {sections.map((section, index) => (
              <motion.button
                key={section.id}
                type="button"
                className={`site-mobile-link${activeSection === section.id ? " is-active" : ""}`}
                onClick={() => handleJump(section.id)}
                initial={{ opacity: 0, x: -12 }}
                animate={{ opacity: 1, x: 0 }}
                exit={{ opacity: 0, x: -12 }}
                transition={{ duration: 0.22, delay: index * 0.04 }}
              >
                {section.label}
              </motion.button>
            ))}
          </motion.div>
        )}
      </AnimatePresence>
    </>
  );
}
