import { useEffect, useRef } from "react";
import { useBackgroundStore } from "../../../store/agent-chat/backgroundStore";
import { useSessionStore } from "../../../store/agent-chat/sessionStore";
import { openSSE, type SSEEvent } from "../../../api/agent-chat/sessions";

const EVENT_ICONS: Record<string, string> = {
  start: "◎",
  tool_call: "⚙",
  tools_used: "✓",
  hint: "◈",
  consolidate: "▣",
};

function formatEvent(event: SSEEvent): string {
  switch (event.type) {
    case "start":
      return `开始分析: ${event.message?.slice(0, 80) ?? ""}`;
    case "tool_call":
      return `调用工具: ${event.detail ?? ""}`;
    case "tools_used":
      return `工具汇总: ${(event.tools ?? []).join(", ")}`;
    case "hint":
      return `Hint: ${event.hint ?? ""}`;
    case "consolidate":
      return `记忆压缩: session ${event.session ?? ""}`;
    default:
      return JSON.stringify(event);
  }
}

export default function BackgroundPanel() {
  const logs = useBackgroundStore((s) => s.logs);
  const clear = useBackgroundStore((s) => s.clear);
  const addLog = useBackgroundStore((s) => s.addLog);
  const sessionId = useSessionStore((s) => s.sessionId);
  const sessionStatus = useSessionStore((s) => s.sessionStatus);
  const bottomRef = useRef<HTMLDivElement>(null);
  const esRef = useRef<EventSource | null>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [logs]);

  useEffect(() => {
    if (sessionStatus !== "active" || !sessionId) return;

    // Close previous connection
    esRef.current?.close();

    addLog(`SSE 连接中: ${sessionId}`);

    const es = openSSE(
      sessionId,
      (event: SSEEvent) => {
        // skip "done" per spec (show greyed, or skip entirely)
        if (event.type === "done") return;
        const icon = EVENT_ICONS[event.type] ?? "·";
        addLog(`${icon} ${formatEvent(event)}`);
      },
      () => {
        addLog("· SSE 连接断开，等待重连...");
      }
    );

    esRef.current = es;

    return () => {
      es.close();
      esRef.current = null;
    };
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [sessionId, sessionStatus]);

  return (
    <div className="ac-panel ac-bg-panel">
      <div className="ac-panel__header">
        <span className="ac-panel__label">BACKGROUND AGENT</span>
        <div style={{ display: "flex", alignItems: "center", gap: "0.75rem" }}>
          {sessionStatus === "active" && (
            <span className="ac-bg-panel__status-dot" title="SSE 已连接" />
          )}
          <button className="ac-panel__action" onClick={clear}>清空</button>
        </div>
      </div>
      <div className="ac-bg-panel__logs">
        {logs.length === 0 ? (
          <p className="ac-bg-panel__empty">等待 Background Agent 事件...</p>
        ) : (
          logs.map((log) => (
            <div key={log.id} className="ac-bg-panel__entry">
              <span className="ac-bg-panel__ts">[{log.timestamp}]</span>
              <span className="ac-bg-panel__text">{log.text}</span>
            </div>
          ))
        )}
        <div ref={bottomRef} />
      </div>
    </div>
  );
}
