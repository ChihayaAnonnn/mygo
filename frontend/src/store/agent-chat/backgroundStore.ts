import { create } from "zustand";

export interface LogEntry {
  id: string;
  timestamp: string;
  text: string;
}

interface BackgroundState {
  logs: LogEntry[];
  addLog: (text: string) => void;
  clear: () => void;
}

let counter = 0;

export const useBackgroundStore = create<BackgroundState>((set) => ({
  logs: [],
  addLog: (text) => {
    const now = new Date();
    const ts = now.toTimeString().slice(0, 8);
    const id = `log-${Date.now()}-${counter++}`;
    set((s) => {
      const next = [...s.logs, { id, timestamp: ts, text }];
      return { logs: next.length > 100 ? next.slice(next.length - 100) : next };
    });
  },
  clear: () => set({ logs: [] }),
}));
