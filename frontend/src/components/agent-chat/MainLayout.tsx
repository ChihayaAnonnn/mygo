import { useEffect } from "react";
import { motion, type Variants } from "framer-motion";
import { useSessionStore } from "../../store/agent-chat/sessionStore";
import { closeSession } from "../../api/agent-chat/sessions";
import { useBackgroundStore } from "../../store/agent-chat/backgroundStore";
import ChatPanel from "./chat/ChatPanel";
import BackgroundPanel from "./background/BackgroundPanel";
import WorkspacePanel from "./workspace/WorkspacePanel";

const panelVariants: Variants = {
  hidden: { opacity: 0, y: 16 },
  visible: (i: number) => ({
    opacity: 1,
    y: 0,
    transition: { delay: i * 0.1, duration: 0.45, ease: [0.16, 1, 0.3, 1] as [number, number, number, number] },
  }),
};

export default function MainLayout() {
  const { username, sessionId, reset } = useSessionStore();
  const addLog = useBackgroundStore((s) => s.addLog);

  // Close session on page unload
  useEffect(() => {
    const handleUnload = () => {
      if (sessionId) {
        navigator.sendBeacon(`${import.meta.env.VITE_AGENT_API_URL ?? "http://localhost:8000"}/sessions/${sessionId}`);
      }
    };
    window.addEventListener("beforeunload", handleUnload);
    return () => window.removeEventListener("beforeunload", handleUnload);
  }, [sessionId]);

  async function handleSettings() {
    if (sessionId) {
      try {
        await closeSession(sessionId);
        addLog(`Session 已关闭: ${sessionId}`);
      } catch {
        // ignore - session may already be gone
      }
    }
    reset();
  }

  return (
    <div className="ac-main">
      <header className="ac-main__header">
        <span className="ac-main__brand">ASYNC AGENT</span>
        <div className="ac-main__header-right">
          <span className="ac-main__user">{username}</span>
          <button className="ac-main__settings" onClick={handleSettings}>
            Settings
          </button>
        </div>
      </header>

      <div className="ac-main__body">
        <motion.div
          className="ac-main__left"
          custom={0}
          initial="hidden"
          animate="visible"
          variants={panelVariants}
        >
          <ChatPanel />
        </motion.div>

        <div className="ac-main__right">
          <motion.div
            className="ac-main__right-top"
            custom={1}
            initial="hidden"
            animate="visible"
            variants={panelVariants}
          >
            <BackgroundPanel />
          </motion.div>

          <motion.div
            className="ac-main__right-bottom"
            custom={2}
            initial="hidden"
            animate="visible"
            variants={panelVariants}
          >
            <WorkspacePanel />
          </motion.div>
        </div>
      </div>
    </div>
  );
}
