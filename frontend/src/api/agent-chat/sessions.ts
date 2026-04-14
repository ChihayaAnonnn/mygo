import { agentClient } from "./client";
import { useSessionStore } from "../../store/agent-chat/sessionStore";

export interface CreateSessionResponse {
  session_id: string;
  workspace: string;
}

export interface SessionMessageResponse {
  session_id: string;
  response: string;
}

export interface ListSessionsResponse {
  sessions: string[];
}

export interface CloseSessionResponse {
  closed: string;
}

export async function createSession(
  user: string,
  sessionId?: string
): Promise<CreateSessionResponse> {
  const { data } = await agentClient.post<CreateSessionResponse>("/sessions", {
    user,
    ...(sessionId ? { session_id: sessionId } : {}),
  });
  return data;
}

export async function sendSessionMessage(
  sessionId: string,
  message: string,
  timeout = 60
): Promise<SessionMessageResponse> {
  const { data } = await agentClient.post<SessionMessageResponse>(
    `/sessions/${sessionId}/messages`,
    { message, timeout }
  );
  return data;
}

export async function listSessions(): Promise<ListSessionsResponse> {
  const { data } = await agentClient.get<ListSessionsResponse>("/sessions");
  return data;
}

export async function closeSession(sessionId: string): Promise<CloseSessionResponse> {
  const { data } = await agentClient.delete<CloseSessionResponse>(
    `/sessions/${sessionId}`
  );
  return data;
}

export type SSEEventType =
  | "start"
  | "tool_call"
  | "tools_used"
  | "hint"
  | "done"
  | "consolidate"
  | "keepalive";

export interface SSEEvent {
  type: SSEEventType;
  session?: string;
  message?: string;
  detail?: string;
  tools?: string[];
  tools_used?: string[];
  hint?: string;
}

const BASE_URL = import.meta.env.VITE_AGENT_API_URL ?? "http://localhost:8000";

export function openSSE(
  sessionId: string,
  onEvent: (event: SSEEvent) => void,
  onError?: () => void
): EventSource {
  const apiKey = useSessionStore.getState().apiKey;
  const url = new URL(`${BASE_URL}/sessions/${sessionId}/events`);
  if (apiKey) url.searchParams.set("api_key", apiKey);

  const es = new EventSource(url.toString());

  es.onmessage = (e) => {
    try {
      const event: SSEEvent = JSON.parse(e.data);
      if (event.type !== "keepalive") {
        onEvent(event);
      }
    } catch {
      // ignore malformed
    }
  };

  es.onerror = () => {
    onError?.();
  };

  return es;
}
