import { createBrowserRouter } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import HomePage from "../pages/HomePage";
import OpenClawRetiredPage from "../pages/OpenClawRetiredPage";
import WorkDetailPage from "../pages/WorkDetailPage";
import NotFoundPage from "../pages/NotFoundPage";
import AgentChatPage from "../pages/AgentChatPage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "openclaw", element: <OpenClawRetiredPage /> },
      { path: "works/:slug", element: <WorkDetailPage /> },
    ],
  },
  { path: "agent-chat", element: <AgentChatPage /> },
  { path: "*", element: <NotFoundPage /> },
]);

