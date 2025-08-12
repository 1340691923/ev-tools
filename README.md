# EV Tools - ElasticView插件

## 📖 项目简介

EV Tools 是一个基于 ElasticView 平台的 Elasticsearch 管理插件，为开发者和管理员提供直观、高效的 Elasticsearch 集群管理工具。支持索引管理、文档操作、集群监控、快照管理等功能。

## ✨ 主要功能

### 🔍 索引管理
- **索引操作**：创建、删除、打开、关闭索引
- **索引设置**：获取和更新索引设置
- **索引映射**：查看和修改索引映射
- **重建索引**：支持索引重建和数据迁移
- **索引优化**：刷新、清空缓存、强制合并等操作

### 🎯 别名管理
- **别名操作**：添加、移除、移动别名
- **批量别名**：支持批量添加多个别名
- **别名查询**：查看索引的别名信息

### 📊 文档操作
- **文档 CRUD**：创建、读取、更新、删除文档
- **批量操作**：支持批量删除文档
- **文档搜索**：基于条件的文档搜索
- **数据导出**：支持查询结果导出

### 🛠️ 集群监控
- **集群健康**：实时监控集群健康状态
- **节点信息**：查看节点详细信息（CPU、内存、磁盘等）
- **分片状态**：监控分片分配和状态
- **索引统计**：查看索引文档数量和存储情况
- **集群统计**：获取集群整体统计信息

### 📦 快照管理
- **快照仓库**：创建、删除、查看快照仓库
- **快照操作**：创建、删除、恢复快照
- **快照状态**：监控快照创建和恢复进度

### 🎛️ 任务管理
- **任务列表**：查看正在执行的任务
- **任务控制**：取消正在执行的任务

## 🛠️ 技术栈

### 后端
- **Go 1.20+**：高性能后端服务
- **Gin**：轻量级 Web 框架
- **Elasticsearch Go Client**：官方 ES Go 客户端
- **Eve Plugin SDK**：ElasticView 插件开发SDK

### 前端
- **Vue 3**：渐进式 JavaScript 框架
- **Element Plus**：Vue 3 组件库
- **Monaco Editor**：代码编辑器
- **TypeScript**：类型安全的 JavaScript

## 📦 安装要求

### 环境要求
- **Go 版本**：>= 1.20
- **Node 版本**：>= 20.14.0
- **ElasticView**：已启动基座程序
- **Elasticsearch**：目标 ES 集群

### 开发工具安装
```bash
# 安装 ElasticView 插件开发工具
go install github.com/1340691923/eve-plugin-sdk-go/cmd/ev_plugin_cli@v0.0.20

# 安装 pnpm
npm install -g pnpm
```

## 🚀 快速开始

### 1. 下载依赖

```bash
# 检查项目依赖和环境（项目根目录运行）
ev_plugin_cli doctor 

# 使用 ev_plugin_cli 下载项目依赖
ev_plugin_cli install

# 或者手动安装依赖
# 后端：go mod tidy
# 前端：cd frontend && pnpm install
```

### 2. 开发模式运行

```bash
ev_plugin_cli dev
```

### 3. 构建插件
```bash
# 使用 ev_plugin_cli 构建插件包（自动构建前端和后端）
ev_plugin_cli build
```

## 📁 项目结构

```
ev-tools/
├── backend/                 # 后端项目目录
│   ├── services/           # 业务服务层
│   │   ├── alias_service/  # 别名管理服务
│   │   ├── cat_service/    # 集群监控服务
│   │   ├── es_backup_service/ # 快照管理服务
│   │   ├── es_doc_service/ # 文档操作服务
│   │   ├── es_service/     # 通用 ES 服务
│   │   ├── es_task_service/ # 任务管理服务
│   │   ├── index_service/  # 索引管理服务
│   │   └── navicat_service/ # 数据库管理服务
│   ├── dto/                # 数据传输对象
│   ├── my_error/           # 自定义异常处理
│   ├── response/           # 响应处理
│   └── router/             # 路由定义
├── frontend/               # 前端项目目录
│   ├── src/
│   │   ├── api/            # API 接口
│   │   ├── components/     # Vue 组件
│   │   ├── views/          # 页面视图
│   │   └── utils/          # 工具函数
│   └── package.json
├── main.go                 # 主程序入口
├── plugin.json             # 插件配置
└── README.md
```

## ⚙️ 配置说明

### plugin.json 配置
```json
{
  "developer": "肖大侠",
  "version": "0.0.20",
  "plugin_name": "ev工具箱",
  "plugin_alias": "ev-tools",
  "frontend_debug": false,
  "frontend_dev_port": 7002,
  "frontend_routes": [
    {
      "path": "indices",
      "name": "indices",
      "meta": {
        "title": "索引管理",
        "icon": "el-icon-coin"
      }
    },
    {
      "path": "reindex",
      "name": "reindex",
      "meta": {
        "title": "重建索引",
        "icon": "el-icon-document-copy"
      }
    },
    {
      "path": "cat",
      "name": "cat",
      "meta": {
        "title": "ES状态",
        "icon": "el-icon-pie-chart"
      }
    },
    {
      "path": "rest",
      "name": "rest",
      "meta": {
        "title": "开发工具",
        "icon": "el-icon-search"
      }
    },
    {
      "path": "task",
      "name": "task",
      "meta": {
        "title": "任务",
        "icon": "el-icon-ticket"
      }
    },
    {
      "path": "back-up",
      "name": "back-up",
      "meta": {
        "title": "快照存储库",
        "icon": "el-icon-first-aid-kit"
      }
    },
    {
      "path": "snapshot",
      "name": "index",
      "meta": {
        "title": "快照管理",
        "icon": "el-icon-search"
      }
    },
    {
      "path": "navicat",
      "name": "navicat",
      "meta": {
        "title": "Navicat",
        "icon": "el-icon-first-aid-kit"
      }
    }
  ]
}
```

## 📝 使用说明

### 索引管理使用
1. **索引浏览**：在左侧查看所有索引列表
2. **索引操作**：
   - 创建索引：设置索引名称和映射
   - 删除索引：删除不需要的索引
   - 索引设置：修改索引设置参数
   - 索引映射：查看和修改字段映射
3. **索引优化**：执行刷新、清空缓存等操作

### 集群监控使用
1. **健康状态**：查看集群整体健康状态
2. **节点信息**：监控各节点的资源使用情况
3. **分片状态**：查看分片分配和状态
4. **索引统计**：查看各索引的文档数量和存储情况

### 文档操作使用
1. **文档查询**：使用 REST API 查询文档
2. **文档编辑**：修改文档内容
3. **批量操作**：批量删除文档
4. **数据导出**：导出查询结果

### 快照管理使用
1. **仓库管理**：创建和管理快照仓库
2. **快照操作**：创建、删除、恢复快照
3. **状态监控**：查看快照创建和恢复进度

## 🔍 功能特性详解

### 索引管理功能
- **可视化操作**：通过界面操作，无需记忆复杂命令
- **批量操作**：支持批量处理多个索引
- **实时反馈**：操作结果实时显示
- **错误处理**：友好的错误提示和处理

### 集群监控功能
- **实时监控**：集群状态实时更新
- **多维度统计**：CPU、内存、磁盘、网络等多维度监控
- **可视化展示**：图表化展示监控数据
- **告警功能**：异常状态告警提示

### 文档操作功能
- **REST API**：支持完整的 REST API 操作
- **语法高亮**：JSON 格式语法高亮
- **智能提示**：API 参数智能提示
- **历史记录**：操作历史记录保存

### 快照管理功能
- **仓库管理**：完整的快照仓库生命周期管理
- **快照操作**：创建、删除、恢复快照
- **进度监控**：实时监控快照操作进度
- **状态查询**：查询快照和恢复状态

## 📄 许可证

本项目采用 MIT 许可证

## 🙏 致谢

- [ElasticView](https://github.com/1340691923/ElasticView) - 提供优秀的插件开发平台
- [Elasticsearch](https://www.elastic.co/) - 强大的搜索引擎
- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [Element Plus](https://element-plus.org/) - Vue 3 组件库

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 微信：qq1340691923

---

⭐ 如果这个项目对您有帮助，请给我们一个Star！

