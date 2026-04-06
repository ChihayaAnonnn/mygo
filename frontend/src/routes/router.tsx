import { createBrowserRouter } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import HomePage from "../pages/HomePage";
import WorkDetailPage from "../pages/WorkDetailPage";
import OpenClawShowcasePage from "../pages/OpenClawShowcasePage";
import NotFoundPage from "../pages/NotFoundPage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "works/:slug", element: <WorkDetailPage /> },
      { path: "openclaw", element: <OpenClawShowcasePage /> },
    ],
  },
  { path: "*", element: <NotFoundPage /> },
]);

