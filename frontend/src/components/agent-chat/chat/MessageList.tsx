import { useEffect, useRef } from "react";
import { useMessageStore } from "../../../store/agent-chat/messageStore";
import MessageBubble from "./MessageBubble";

export default function MessageList() {
  const messages = useMessageStore((s) => s.messages);
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  if (messages.length === 0) {
    return (
      <div className="ac-chat__empty">
        <p className="ac-chat__empty-title">Async Agent</p>
        <p className="ac-chat__empty-hint">发送消息开始对话</p>
      </div>
    );
  }

  return (
    <div className="ac-message-list">
      {messages.map((msg) => (
        <MessageBubble key={msg.id} message={msg} />
      ))}
      <div ref={bottomRef} />
    </div>
  );
}
