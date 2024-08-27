import pluginJson from '../../../plugin.json'

export const routes = [
  {
    path: `/${pluginJson.plugin_alias}`,//根目录必须为插件名
    component: ()=>import("../layouts/Layout.vue"),
    children: [
      {
        path: 'indices',
        component: ()=>import('../views/indices/index.vue'),
        name: 'indices',
        meta: {
          title: '索引管理',
          icon: 'el-icon-coin'
        }
      },
      {
        path: 'reindex',
        component: ()=>import('../views/indices/reindex.vue'),
        name: 'reindex',
        meta: {
          title: '重建索引',
          icon: 'el-icon-document-copy'
        }
      },
      {
        path: 'cat',
        component: ()=>import('../views/cat/index.vue'),
        name: 'cat',
        meta: {
          title: 'ES状态',
          icon: 'el-icon-pie-chart'
        }
      },
      {
        path: 'rest',
        component: ()=>import('../views/rest/index.vue'),
        name: 'rest',
        meta: {
          title: '开发工具',
          icon: 'el-icon-search'
        }
      },
      {
        path: 'task',
        component: ()=>import('../views/task/index.vue'),
        name: 'task',
        meta: {
          title: '任务',
          icon: 'el-icon-notebook-2'
        }
      },
      {
        path: 'back-up',
        component: ()=>import('../views/back-up/index.vue'),
        name: 'back-up',
        meta: {
          title: '快照存储库',
          icon: 'el-icon-first-aid-kit'
        }
      },
      {
        path: 'snapshot',
        component: ()=>import('../views/back-up/snapshot.vue'),
        name: 'index',
        meta: {
          title: '快照管理',
          icon: "el-icon-search"
        }
      },
      {
        path: 'navicat',
        component: ()=>import('../views/navicat/index.vue'),
        name: 'navicat',
        meta: {
          title: 'Navicat',
          icon: 'el-icon-first-aid-kit'
        }
      }
    ]
  },
]

export default routes
