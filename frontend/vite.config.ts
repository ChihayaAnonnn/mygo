import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { fileURLToPath, URL } from 'node:url';

export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 5173,
    fs: {
      allow: ["/workspace/mygo/frontend", "/workspace/data/posts"],
    },
  },
  resolve: {
    alias: {
      '@posts': fileURLToPath(new URL('./posts', import.meta.url)),
    },
  },
});
