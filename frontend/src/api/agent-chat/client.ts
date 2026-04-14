import axios from "axios";
import { useSessionStore } from "../../store/agent-chat/sessionStore";

const BASE_URL = import.meta.env.VITE_AGENT_API_URL ?? "http://localhost:8000";

export const agentClient = axios.create({ baseURL: BASE_URL });

agentClient.interceptors.request.use((config) => {
  const apiKey = useSessionStore.getState().apiKey;
  if (apiKey) {
    config.headers["X-Dashscope-Key"] = apiKey;
  }
  return config;
});
