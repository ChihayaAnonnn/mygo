import React, { useEffect, useRef, useState } from "react";
import { io } from "socket.io-client";
import { ArrowLeft, RefreshCw } from "lucide-react";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";

// 类型定义
type LobsterAction =
  | "idle"
  | "wave"
  | "dance"
  | "sleep"
  | "work"
  | "eat"
  | "play"
  | "react";

type LobsterEmotion = "happy" | "curious" | "sleepy" | "excited" | "calm" | "alert";

interface OpenClawStatus {
  status: "online" | "busy" | "away" | "sleeping";
  activity: string;
  mood: string;
  message: string;
  timestamp: number;
}

interface SocketAvatarAction {
  action: string;
  emotion: string;
  message: string;
}

const ACTION_LABELS: Record<LobsterAction, string> = {
  idle: "待机",
  wave: "挥钳问好",
  dance: "跳舞",
  sleep: "休息",
  work: "工作中",
  eat: "吃东西",
  play: "玩耍",
  react: "回应",
};

const EMOTION_LABELS: Record<LobsterEmotion, string> = {
  happy: "开心",
  curious: "好奇",
  sleepy: "困倦",
  excited: "兴奋",
  calm: "平静",
  alert: "警觉",
};

const toLobsterAction = (value: string): LobsterAction => {
  const validActions: LobsterAction[] = [
    "idle",
    "wave",
    "dance",
    "sleep",
    "work",
    "eat",
    "play",
    "react",
  ];

  return validActions.includes(value as LobsterAction)
    ? (value as LobsterAction)
    : "idle";
};

const toLobsterEmotion = (value: string): LobsterEmotion => {
  const validEmotions: LobsterEmotion[] = [
    "happy",
    "curious",
    "sleepy",
    "excited",
    "calm",
    "alert",
  ];

  return validEmotions.includes(value as LobsterEmotion)
    ? (value as LobsterEmotion)
    : "happy";
};

function OpenClawShowcasePage() {
  const [status, setStatus] = useState<OpenClawStatus>({
    status: "online",
    activity: "初始化中...",
    mood: "curious",
    message: "你好，我是 OpenClaw。这里是我作为数字伙伴的展示页。",
    timestamp: Date.now(),
  });
  const [connected, setConnected] = useState(false);
  const [avatarAction, setAvatarAction] = useState<LobsterAction>("idle");
  const [avatarEmotion, setAvatarEmotion] = useState<LobsterEmotion>("happy");
  const [updatedAt, setUpdatedAt] = useState(Date.now());
  const [toastMessage, setToastMessage] = useState<string | null>(null);
  const toastTimerRef = useRef<number | null>(null);

  const showToast = (message: string) => {
    setToastMessage(message);

    if (toastTimerRef.current) {
      window.clearTimeout(toastTimerRef.current);
    }

    toastTimerRef.current = window.setTimeout(() => {
      setToastMessage(null);
    }, 2400);
  };

  useEffect(() => {
    const newSocket = io("/ws/openclaw", {
      transports: ["websocket"],
      reconnection: true,
      reconnectionAttempts: 5,
      reconnectionDelay: 1000,
    });

    newSocket.on("connect", () => {
      setConnected(true);
      newSocket.emit("subscribe", { channel: "status" });
      newSocket.emit("subscribe", { channel: "avatar" });
      newSocket.emit("subscribe", { channel: "info" });
      newSocket.emit("subscribe", { channel: "notifications" });
    });

    newSocket.on("disconnect", () => {
      setConnected(false);
    });

    newSocket.on("status_update", (data: OpenClawStatus) => {
      setStatus(data);
    });

    newSocket.on("avatar_action", (data: SocketAvatarAction) => {
      setAvatarAction(toLobsterAction(data.action));
      setAvatarEmotion(toLobsterEmotion(data.emotion));
      if (data.message) {
        setStatus((currentStatus) => ({
          ...currentStatus,
          message: data.message,
          timestamp: Date.now(),
        }));
      }
    });

    newSocket.on("info_update", () => {
      setUpdatedAt(Date.now());
    });

    newSocket.on("notification", (data: { message: string }) => {
      showToast(data.message);
    });

    return () => {
      if (toastTimerRef.current) {
        window.clearTimeout(toastTimerRef.current);
      }
      newSocket.disconnect();
    };
  }, []);

  const handleRefresh = () => {
    setUpdatedAt(Date.now());
    showToast("OpenClaw 内容已刷新。");
  };

  return (
    <div className="openclaw-page">
      {toastMessage ? (
        <div className="openclaw-toast" role="status" aria-live="polite">
          {toastMessage}
        </div>
      ) : null}

      <motion.section
        className="content-section openclaw-overview"
        initial={{ opacity: 0, y: 24 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.55 }}
      >
        <div className="openclaw-page-topbar">
          <Link to="/" className="inline-link">
            <ArrowLeft size={16} />
            返回首页
          </Link>

          <button
            type="button"
            className="openclaw-refresh-button"
            onClick={handleRefresh}
          >
            <RefreshCw size={16} />
            刷新内容
          </button>
        </div>

        <div className="openclaw-hero-grid">
          <div className="section-headline openclaw-copy">
            <p className="section-kicker">Work / Character</p>
            <h1 id="openclaw-title" className="openclaw-page-title">
              OpenClaw Showcase
            </h1>
            <p className="openclaw-description">
              这是一个把数字伙伴做成作品页的展示入口。它不再伪装成一块复杂控制台，而是用更清晰的结构呈现角色状态、少量互动和信息摘要。
            </p>

            <div className="openclaw-meta-row">
              <span className={`openclaw-pill${connected ? " is-live" : ""}`}>
                <span className="openclaw-pill-dot" aria-hidden="true" />
                {connected ? "Live connection" : "Local preview"}
              </span>
              <span className="openclaw-pill">{status.activity}</span>
              <span className="openclaw-pill">
                {EMOTION_LABELS[avatarEmotion]} / {ACTION_LABELS[avatarAction]}
              </span>
            </div>

            <p className="openclaw-note">{status.message}</p>
          </div>

          <div className="openclaw-stage">
            <div
              className="openclaw-avatar-frame openclaw-placeholder-frame"
              role="img"
              aria-label="OpenClaw visual placeholder"
            >
              <div className="openclaw-placeholder-surface" aria-hidden="true" />
            </div>
            <div className="openclaw-stage-caption">
              <span>{EMOTION_LABELS[avatarEmotion]}</span>
              <span>{ACTION_LABELS[avatarAction]}</span>
              <span>{new Date(updatedAt).toLocaleTimeString()}</span>
            </div>
          </div>
        </div>
      </motion.section>

    </div>
  );
}

export default OpenClawShowcasePage;