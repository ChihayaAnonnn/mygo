import { agentClient } from "./client";

export interface InitResponse {
  workspace: string;
  created: string[];
}

export interface ChatResponse {
  response: string;
  session_id: string;
}

export interface HeartbeatResponse {
  result: string;
}

export async function initWorkspace(user: string): Promise<InitResponse> {
  const { data } = await agentClient.post<InitResponse>("/workspace/init", null, {
    params: { user },
  });
  return data;
}

export async function sendMessage(
  message: string,
  user: string,
  sessionId?: string
): Promise<ChatResponse> {
  const { data } = await agentClient.post<ChatResponse>("/chat", {
    message,
    user,
    ...(sessionId ? { session_id: sessionId } : {}),
  });
  return data;
}

export async function triggerHeartbeat(user: string): Promise<HeartbeatResponse> {
  const { data } = await agentClient.post<HeartbeatResponse>("/heartbeat", null, {
    params: { user },
  });
  return data;
}
