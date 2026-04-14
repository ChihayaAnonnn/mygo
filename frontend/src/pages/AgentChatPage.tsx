import { useEffect, useState } from "react";
import { useSessionStore } from "../store/agent-chat/sessionStore";
import { listSessions } from "../api/agent-chat/sessions";
import { useBackgroundStore } from "../store/agent-chat/backgroundStore";
import SetupScreen from "../components/agent-chat/SetupScreen";
import MainLayout from "../components/agent-chat/MainLayout";
import "../styles/agent-chat.css";

export default function AgentChatPage() {
  const { sessionId, username, apiKey, sessionStatus, setSessionStatus, reset } =
    useSessionStore();
  const addLog = useBackgroundStore((s) => s.addLog);
  const [checking, setChecking] = useState(true);

  useEffect(() => {
    async function verifySession() {
      // No stored credentials → show setup
      if (!username || !apiKey) {
        setChecking(false);
        return;
      }

      // Credentials exist but no session ID → show setup (will re-init)
      if (!sessionId) {
        setChecking(false);
        return;
      }

      // Have session ID → verify it's still alive
      try {
        const { sessions } = await listSessions();
        if (sessions.includes(sessionId)) {
          setSessionStatus("active");
          addLog(`Session 已恢复: ${sessionId}`);
        } else {
          // Session expired on backend → need fresh init
          addLog(`Session 已过期，请重新配置`);
          reset();
        }
      } catch {
        // Backend unreachable → still allow UI to show but mark idle
        setSessionStatus("idle");
      } finally {
        setChecking(false);
      }
    }

    verifySession();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (checking) {
    return (
      <div className="ac-setup">
        <p style={{ fontFamily: "var(--ac-mono)", fontSize: "0.8rem", color: "var(--ac-text2)" }}>
          验证 Session...
        </p>
      </div>
    );
  }

  const isReady = !!(username && apiKey && sessionId && sessionStatus === "active");
  return isReady ? <MainLayout /> : <SetupScreen />;
}
