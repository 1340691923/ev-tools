import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import qiankun from 'vite-plugin-qiankun'
import pluginJson from './../plugin.json';
import { resolve } from "path";
import { dynamicBase } from 'vite-plugin-dynamic-base'
import monacoEditorPlugin from "vite-plugin-monaco-editor";

const pathSrc = resolve(__dirname, "src");

export default defineConfig({
  resolve: {
    alias: {
      "@": pathSrc,
    },
  },

  base: process.env.NODE_ENV === "production" ? "/__dynamic_base__/" : "/",

  isProduction:false,
  plugins: [
    vue(),
    monacoEditorPlugin.default({
      languageWorkers: ['json'],
    }),
    qiankun(pluginJson.plugin_alias, {
      useDevMode: true
    }),
    dynamicBase({
      publicPath: 'window.proxy.__INJECTED_PUBLIC_PATH_BY_QIANKUN__.substring(0,window.proxy.__INJECTED_PUBLIC_PATH_BY_QIANKUN__.length-1)',
      transformIndexHtml:  true
    }),

  ],

  server: {
    port: pluginJson['frontend_dev_port'],
    headers: {
      'Access-Control-Allow-Origin': '*'
    }
  }
})
