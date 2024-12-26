import {setupEvPlugin} from './plugin_sdk/setup'

import JsonViewer from "vue3-json-viewer";
import "vue3-json-viewer/dist/index.css";

import JsonExcel from "vue-json-excel3";


// 插件启动所需安装第三方插件通过此方法
setupEvPlugin((app)=>{
  app.component("downloadExcel", JsonExcel);

  app.use(JsonViewer)

})
