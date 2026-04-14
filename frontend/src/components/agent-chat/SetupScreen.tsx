import { useState } from "react";
import { motion } from "framer-motion";
import { useSessionStore } from "../../store/agent-chat/sessionStore";
import { initWorkspace } from "../../api/agent-chat/chat";
import { createSession } from "../../api/agent-chat/sessions";
import { useBackgroundStore } from "../../store/agent-chat/backgroundStore";

export default function SetupScreen() {
  const [username, setUsername] = useState("");
  const [apiKey, setApiKey] = useState("sk-b0f2f50007314aa5b0a386d0b7722369");
  const [loading, setLoading] = useState(false);
  const [step, setStep] = useState<"" | "workspace" | "session">("");
  const [error, setError] = useState("");

  const setCredentials = useSessionStore((s) => s.setCredentials);
  const setSession = useSessionStore((s) => s.setSession);
  const addLog = useBackgroundStore((s) => s.addLog);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!username.trim() || !apiKey.trim()) {
      setError("用户名和 API Key 均为必填项");
      return;
    }
    setLoading(true);
    setError("");

    // Store credentials first so axios interceptor can use them
    setCredentials(username.trim(), apiKey.trim());

    try {
      // Step 1: initialize workspace
      setStep("workspace");
      await initWorkspace(username.trim());
      addLog(`工作区初始化完成，用户: ${username.trim()}`);

      // Step 2: create persistent session
      setStep("session");
      const res = await createSession(username.trim());
      setSession(res.session_id, "active");
      addLog(`Session 已创建: ${res.session_id}`);
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : "初始化失败";
      setError(`初始化失败：${msg}`);
      // clear stored credentials on failure
      useSessionStore.getState().reset();
    } finally {
      setLoading(false);
      setStep("");
    }
  }

  const stepLabel = step === "workspace"
    ? "初始化工作区..."
    : step === "session"
    ? "启动 Agent Session..."
    : "开始使用";

  return (
    <div className="ac-setup">
      <motion.div
        className="ac-setup__card"
        initial={{ opacity: 0, y: 32 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5, ease: [0.16, 1, 0.3, 1] as [number, number, number, number] }}
      >
        <div className="ac-setup__header">
          <span className="ac-setup__eyebrow">ASYNC AGENT</span>
          <h1 className="ac-setup__title">配置工作区</h1>
          <p className="ac-setup__desc">输入用户名与 Dashscope API Key 以启动 Agent 会话。</p>
        </div>

        <form className="ac-setup__form" onSubmit={handleSubmit}>
          <div className="ac-field">
            <label className="ac-field__label" htmlFor="username">用户名</label>
            <input
              id="username"
              className="ac-field__input"
              type="text"
              placeholder="alice"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              autoComplete="username"
              autoFocus
            />
          </div>
          <div className="ac-field">
            <label className="ac-field__label" htmlFor="apiKey">Dashscope API Key</label>
            <input
              id="apiKey"
              className="ac-field__input"
              type="password"
              placeholder="sk-xxxxxxxxxxxxxxxx"
              value={apiKey}
              onChange={(e) => setApiKey(e.target.value)}
              autoComplete="current-password"
            />
          </div>
          {error && <p className="ac-setup__error">{error}</p>}
          <button className="ac-setup__btn" type="submit" disabled={loading}>
            {loading ? stepLabel : "开始使用"}
          </button>
        </form>
      </motion.div>
    </div>
  );
}
