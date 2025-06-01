package main

import (
	"embed"
	_ "embed"
	"ev-plugin/backend/migrate"
	"ev-plugin/backend/router"
	"ev-plugin/frontend"
	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

//go:embed plugin.json
var pluginJsonBytes []byte

//go:embed logo.png
var logoPng embed.FS

func main() {
	flag.Parse()

	plugin_server.Serve(plugin_server.ServeOpts{
		Assets: &plugin_server.Assets{
			PluginJsonBytes: pluginJsonBytes,
			FrontendFiles:   frontend.StatisFs,
			Icon:            logoPng,
		},
		Migration: &build.Gormigrate{Migrations: []*build.Migration{
			migrate.V0_0_1(),
		}}, //数据版本迁移
		RegisterRoutes: router.NewRouter, //后端路由
	})
}
