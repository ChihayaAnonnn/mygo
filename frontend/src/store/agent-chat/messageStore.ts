import { create } from "zustand";

export type MessageRole = "user" | "agent" | "tool" | "error" | "loading";

export interface Message {
  id: string;
  role: MessageRole;
  content: string;
  toolName?: string;
  toolArgs?: unknown;
  timestamp: number;
}

interface MessageState {
  messages: Message[];
  addMessage: (msg: Omit<Message, "id" | "timestamp">) => string;
  updateMessage: (id: string, patch: Partial<Message>) => void;
  removeMessage: (id: string) => void;
  clear: () => void;
}

let counter = 0;

export const useMessageStore = create<MessageState>((set) => ({
  messages: [],
  addMessage: (msg) => {
    const id = `msg-${Date.now()}-${counter++}`;
    set((s) => ({
      messages: [...s.messages, { ...msg, id, timestamp: Date.now() }],
    }));
    return id;
  },
  updateMessage: (id, patch) =>
    set((s) => ({
      messages: s.messages.map((m) => (m.id === id ? { ...m, ...patch } : m)),
    })),
  removeMessage: (id) =>
    set((s) => ({ messages: s.messages.filter((m) => m.id !== id) })),
  clear: () => set({ messages: [] }),
}));
