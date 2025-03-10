package main

import (
	"context"
	_ "embed"
	"ev-plugin/backend/migrate"
	"ev-plugin/backend/router"
	"ev-plugin/frontend"
	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

//go:embed plugin.json
var pluginJsonBytes []byte

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())

	plugin_server.Serve(plugin_server.ServeOpts{
		PluginJsonBytes: pluginJsonBytes, //插件配置
		Migration: &build.Gormigrate{Migrations: []*build.Migration{
			migrate.V0_0_1(),
		}},                                   //数据版本迁移
		FrontendFiles: frontend.StatisFs,     //前端资源
		WebEngine:     router.NewRouter(ctx), //后端路由
		ExitCallback: func() {
			logger.DefaultLogger.Debug("ev工具箱插件退出")
			cancel()
		},
	})
}
