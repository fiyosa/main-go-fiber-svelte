import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vite.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), svelte()],

  publicDir: "../public",
  envDir: "..",

  server: {
    port: 3000,
    host: "localhost",
  },

  build: {
    emptyOutDir: true,
    manifest: false,
    copyPublicDir: false,
    outDir: "../public/build",
    rolldownOptions: {
      input: "./src/main.ts",
      output: {
        codeSplitting: true,
        entryFileNames: "main.js",
        assetFileNames: (assetInfo) => {
          if (assetInfo.names.includes("main.css")) return "main.css";
          return "[hash].[ext]";
        },
      },
    },
  },
});
