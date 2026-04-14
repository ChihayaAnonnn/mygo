import MessageList from "./MessageList";
import ChatInput from "./ChatInput";

export default function ChatPanel() {
  return (
    <div className="ac-panel ac-chat">
      <div className="ac-panel__header">
        <span className="ac-panel__label">CHAT</span>
      </div>
      <div className="ac-chat__body">
        <MessageList />
      </div>
      <ChatInput />
    </div>
  );
}
