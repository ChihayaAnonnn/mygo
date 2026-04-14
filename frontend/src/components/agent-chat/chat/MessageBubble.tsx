import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import type { Message } from "../../../store/agent-chat/messageStore";

interface Props {
  message: Message;
}

export default function MessageBubble({ message }: Props) {
  const { role, content, toolName, toolArgs } = message;

  if (role === "loading") {
    return (
      <div className="ac-bubble ac-bubble--loading">
        <span className="ac-bubble__dot" />
        <span className="ac-bubble__dot" />
        <span className="ac-bubble__dot" />
      </div>
    );
  }

  if (role === "user") {
    return (
      <div className="ac-bubble ac-bubble--user">
        <p className="ac-bubble__text">{content}</p>
      </div>
    );
  }

  if (role === "tool") {
    return (
      <details className="ac-bubble ac-bubble--tool">
        <summary className="ac-bubble__tool-summary">
          <span className="ac-bubble__tool-icon">⚙</span>
          <span className="ac-bubble__tool-name">{toolName ?? "tool_call"}</span>
        </summary>
        {toolArgs !== undefined && (
          <pre className="ac-bubble__tool-args">
            {JSON.stringify(toolArgs, null, 2)}
          </pre>
        )}
        {content && <p className="ac-bubble__tool-result">{content}</p>}
      </details>
    );
  }

  if (role === "error") {
    return (
      <div className="ac-bubble ac-bubble--error">
        <p className="ac-bubble__text">{content}</p>
      </div>
    );
  }

  return (
    <div className="ac-bubble ac-bubble--agent">
      <div className="ac-bubble__markdown"><ReactMarkdown remarkPlugins={[remarkGfm]}>{content}</ReactMarkdown></div>
    </div>
  );
}
