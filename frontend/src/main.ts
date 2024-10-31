import {setupEvPlugin} from './plugin_sdk/setup'

// 引入依赖
import ElementPlus from 'element-plus'
// 引入全局 CSS 样式
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import JsonViewer from "vue3-json-viewer";
import "vue3-json-viewer/dist/index.css";

import JsonExcel from "vue-json-excel3";

import floatBtnDirectives from "./utils/float-btn"

// 插件启动所需安装第三方插件通过此方法
const registerPlugin = (app)=>{
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }

  app.component("downloadExcel", JsonExcel);

  app.use(JsonViewer).use(ElementPlus).use(floatBtnDirectives)
  return app
}




setupEvPlugin(registerPlugin)
