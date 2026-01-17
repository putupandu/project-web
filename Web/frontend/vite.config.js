// frontend/vite.config.js
import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  return {
    plugins: [react()],
    server: {
      host: true, // bisa diakses LAN
      port: 5173,
      strictPort: true,
      proxy: {
        '/api': {
          target: env.VITE_API_TARGET,
          changeOrigin: true,
          secure: false,
        },
        '/uploads': {
          target: env.VITE_API_TARGET,
          changeOrigin: true,
          secure: false,
        },
      },
    },
  };
});
