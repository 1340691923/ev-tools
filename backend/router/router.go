package router

import (
	"ev-tools/backend/api"
	"ev-tools/backend/response"
	"ev-tools/backend/services/cat_service"
	"ev-tools/backend/services/cluser_settings_service"
	"ev-tools/backend/services/es_backup"
	"ev-tools/backend/services/es_doc_service"
	"ev-tools/backend/services/es_service"
	"ev-tools/backend/services/es_task_service"
	"ev-tools/backend/services/index_service"
	"ev-tools/backend/services/navicat_service"
	"ev-tools/frontend"
	"github.com/1340691923/eve-plugin-sdk-go/backend"
	"github.com/1340691923/eve-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	engine              *gin.Engine
	esBackUpController  *api.EsBackUpController
	esController        *api.EsController
	esCrudController    *api.EsCrudController
	esDocController     *api.EsDocController
	esIndexController   *api.EsIndexController
	esMappingController *api.EsMappingController
	esTaskController    *api.EsTaskController
}

// 依赖注入
func NewWebServer(app *gin.Engine) *WebServer {
	baseController := api.NewBaseController(response.NewResponse())
	esBackUpController := api.NewEsBackUpController(baseController, es_backup.NewEsBackUpService(cluser_settings_service.NewClusterSettingsService()))
	esController := api.NewEsController(baseController, cat_service.NewCatService(), es_service.NewEsService())
	esCurdController := api.NewEsCrudController(baseController, navicat_service.NewNavicatService())
	esDocController := api.NewEsDocController(baseController, es_doc_service.NewEsDocService())
	esIndexController := api.NewEsIndexController(baseController, index_service.NewIndexService())
	esMapController := api.NewEsMappingController(baseController, index_service.NewIndexService())
	esTaskController := api.NewEsTaskController(baseController, es_task_service.NewEsTaskService())

	return &WebServer{
		engine:              app,
		esBackUpController:  esBackUpController,
		esController:        esController,
		esCrudController:    esCurdController,
		esDocController:     esDocController,
		esIndexController:   esIndexController,
		esMappingController: esMapController,
		esTaskController:    esTaskController,
	}
}

// 实现web资源接口（webapi） 可用任何实现http.Handle接口的Web框架开发 我这里用gin为例
func NewResourceHandler(app *gin.Engine) backend.CallResourceHandler {

	//前端页面
	//因为前端所用技术可以进行热更新，所以可进行脱离插件控制
	app.Use(frontend.Serve("/", frontend.EmbedFolder(frontend.StatisFs, "dist")))

	//后端api
	webSvr := NewWebServer(app)
	webSvr.runDslHistory()
	webSvr.runEs()
	webSvr.runEsIndex()
	webSvr.runEsTask()
	webSvr.runEsBackUp()
	webSvr.runEsCrud()
	webSvr.runEsDoc()
	webSvr.runEsMap()
	return httpadapter.New(app.Handler())
}
