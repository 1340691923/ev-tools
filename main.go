package main

import (
	_ "embed"
	"encoding/json"
	"ev-tools/backend/migrate"
	"ev-tools/backend/router"
	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
	"github.com/1340691923/eve-plugin-sdk-go/build"
	"github.com/gin-gonic/gin"
)

//go:embed plugin.json
var pluginJsonBytes []byte
var (
	evRpcPort string
	evRpcKey  string
	debug     bool
)

func init() {
	flag.StringVar(&evRpcPort, "evRpcPort", "", "ev基座内网访问端口")
	flag.StringVar(&evRpcKey, "evRpcKey", "", "ev基座访问密钥")
	flag.BoolVar(&debug, "debug", debug, "是否开启调试")
	flag.Parse()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	pluginJson := new(build.PluginJsonData)
	err := json.Unmarshal(pluginJsonBytes, &pluginJson)
	if err != nil {
		panic(err)
	}
	pluginJson.BackendDebug = debug
	plugin_server.Serve(plugin_server.ServeOpts{
		EvRpcPort:           evRpcPort,
		EvRpcKey:            evRpcKey,
		PluginJson:          pluginJson,
		Migration:           migrate.GetMigrates(),
		CallResourceHandler: router.NewResourceHandler(gin.New()),
		ExitCallback: func() {
			logger.DefaultLogger.Debug("进程退出")
		},
	})
}
