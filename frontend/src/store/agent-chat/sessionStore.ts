import { create } from "zustand";
import { persist } from "zustand/middleware";

export type SessionStatus = "idle" | "initializing" | "active";

interface SessionState {
  username: string;
  apiKey: string;
  sessionId: string;
  sessionStatus: SessionStatus;
  setCredentials: (username: string, apiKey: string) => void;
  setSession: (sessionId: string, status: SessionStatus) => void;
  setSessionStatus: (status: SessionStatus) => void;
  reset: () => void;
}

export const useSessionStore = create<SessionState>()(
  persist(
    (set) => ({
      username: "",
      apiKey: "",
      sessionId: "",
      sessionStatus: "idle",
      setCredentials: (username, apiKey) =>
        set({ username, apiKey }),
      setSession: (sessionId, status) =>
        set({ sessionId, sessionStatus: status }),
      setSessionStatus: (sessionStatus) => set({ sessionStatus }),
      reset: () =>
        set({ username: "", apiKey: "", sessionId: "", sessionStatus: "idle" }),
    }),
    {
      name: "agent-chat-session",
      partialize: (s) => ({
        username: s.username,
        apiKey: s.apiKey,
        sessionId: s.sessionId,
      }),
    }
  )
);
