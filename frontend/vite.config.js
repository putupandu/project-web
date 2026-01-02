// frontend/vite.config.js
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0', // WAJIB supaya bisa diakses teman kamu
    port: 5173,
    strictPort: true, // supaya tidak pindah port
    proxy: {
      '/api': {
        target: 'http://172.22.4.199:8080', // gunakan IP LAN kamu
        changeOrigin: true,
        secure: false,
      }
    }
  }
});
