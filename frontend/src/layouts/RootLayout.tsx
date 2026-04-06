import { Outlet } from "react-router-dom";

export default function RootLayout() {
  return (
    <div className="site-shell">
      <main className="site-main">
        <Outlet />
      </main>
    </div>
  );
}
