import { useEffect, useState } from "react";
import { motion } from "framer-motion";

interface LoadingProps {
  isVisible: boolean;
}

const WORDMARK = "mygo";

export default function Loading({ isVisible }: LoadingProps) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [showCursor, setShowCursor] = useState(true);

  useEffect(() => {
    if (!isVisible) {
      return;
    }

    setCurrentIndex(0);
    setShowCursor(true);

    let index = 0;
    let timeoutId: ReturnType<typeof setTimeout> | undefined;

    const type = () => {
      if (index < WORDMARK.length) {
        setCurrentIndex(index + 1);
        index += 1;
        timeoutId = setTimeout(type, 160);
      }
    };

    type();

    const cursorInterval = window.setInterval(() => {
      setShowCursor((value) => !value);
    }, 450);

    return () => {
      window.clearInterval(cursorInterval);
      if (timeoutId) {
        window.clearTimeout(timeoutId);
      }
    };
  }, [isVisible]);

  if (!isVisible) {
    return null;
  }

  return (
    <div className="loading-screen" aria-hidden="true">
      <div className="loading-wordmark">
        {WORDMARK.split("").map((char, index) => {
          if (index >= currentIndex) {
            return null;
          }

          if (index === currentIndex - 1) {
            return (
              <motion.span
                key={`${char}-${index}`}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.18, ease: "easeOut" }}
              >
                {char}
              </motion.span>
            );
          }

          return <span key={`${char}-${index}`}>{char}</span>;
        })}
        <span className={`loading-cursor${showCursor ? " is-visible" : ""}`}>|</span>
      </div>
    </div>
  );
}
