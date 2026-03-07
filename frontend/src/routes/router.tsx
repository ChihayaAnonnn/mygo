import { createBrowserRouter } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import HomePage from "../pages/HomePage";
import PostsPage from "../pages/PostsPage";
import PostDetailPage from "../pages/PostDetailPage";
import AboutPage from "../pages/AboutPage";
import UploadPage from "../pages/UploadPage";
import NotFoundPage from "../pages/NotFoundPage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "posts", element: <PostsPage /> },
      { path: "posts/:slug", element: <PostDetailPage /> },
      { path: "about", element: <AboutPage /> },
    ],
  },
  { path: "upload", element: <UploadPage /> },
  { path: "*", element: <NotFoundPage /> },
]);

