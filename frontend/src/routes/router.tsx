import { createBrowserRouter } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import ComingSoonPage from "../pages/ComingSoonPage";
import AuthPage from "../pages/AuthPage";
import NotFoundPage from "../pages/NotFoundPage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <ComingSoonPage /> },
      { path: "auth", element: <AuthPage /> },
      { path: "login", element: <AuthPage /> },
      { path: "register", element: <AuthPage /> },
    ],
  },
  { path: "*", element: <NotFoundPage /> },
]);

