import { useRef, useState } from "react";
import { useMessageStore } from "../../../store/agent-chat/messageStore";
import { useSessionStore } from "../../../store/agent-chat/sessionStore";
import { useBackgroundStore } from "../../../store/agent-chat/backgroundStore";
import { sendSessionMessage } from "../../../api/agent-chat/sessions";

export default function ChatInput() {
  const [text, setText] = useState("");
  const [sending, setSending] = useState(false);
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const addMessage = useMessageStore((s) => s.addMessage);
  const removeMessage = useMessageStore((s) => s.removeMessage);
  const clear = useMessageStore((s) => s.clear);
  const sessionId = useSessionStore((s) => s.sessionId);
  const addLog = useBackgroundStore((s) => s.addLog);

  async function submit() {
    const trimmed = text.trim();
    if (!trimmed || sending || !sessionId) return;

    setText("");
    if (textareaRef.current) {
      textareaRef.current.style.height = "auto";
    }

    addMessage({ role: "user", content: trimmed });
    const loadingId = addMessage({ role: "loading", content: "" });
    setSending(true);
    addLog(`· 发送消息: ${trimmed.slice(0, 60)}${trimmed.length > 60 ? "..." : ""}`);

    try {
      const res = await sendSessionMessage(sessionId, trimmed);
      removeMessage(loadingId);
      addMessage({ role: "agent", content: res.response });
      addLog(`· Agent 响应完成`);
    } catch (err: unknown) {
      removeMessage(loadingId);
      const msg = err instanceof Error ? err.message : "请求失败";
      addMessage({ role: "error", content: `错误: ${msg}` });
      addLog(`· 请求失败: ${msg}`);
    } finally {
      setSending(false);
    }
  }

  function handleKeyDown(e: React.KeyboardEvent<HTMLTextAreaElement>) {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      submit();
    }
  }

  function handleInput(e: React.ChangeEvent<HTMLTextAreaElement>) {
    setText(e.target.value);
    const el = e.target;
    el.style.height = "auto";
    el.style.height = Math.min(el.scrollHeight, 120) + "px";
  }

  return (
    <div className="ac-chat-input">
      <div className="ac-chat-input__row">
        <textarea
          ref={textareaRef}
          className="ac-chat-input__textarea"
          placeholder="输入消息... (Enter 发送，Shift+Enter 换行)"
          value={text}
          onChange={handleInput}
          onKeyDown={handleKeyDown}
          disabled={sending}
          rows={1}
        />
        <button
          className="ac-chat-input__send"
          onClick={submit}
          disabled={sending || !text.trim() || !sessionId}
          aria-label="发送"
        >
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round">
            <line x1="22" y1="2" x2="11" y2="13" />
            <polygon points="22 2 15 22 11 13 2 9 22 2" />
          </svg>
        </button>
      </div>
      <div className="ac-chat-input__actions">
        <button
          className="ac-chat-input__clear"
          onClick={() => clear()}
          disabled={sending}
        >
          清空会话
        </button>
      </div>
    </div>
  );
}
