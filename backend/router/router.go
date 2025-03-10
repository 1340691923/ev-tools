package router

import (
	"context"
	"ev-plugin/backend/api"
	"ev-plugin/backend/response"
	"ev-plugin/backend/services/cat_service"
	"ev-plugin/backend/services/cluser_settings_service"
	"ev-plugin/backend/services/es_backup_service"
	"ev-plugin/backend/services/es_doc_service"
	"ev-plugin/backend/services/es_service"
	"ev-plugin/backend/services/es_task_service"
	"ev-plugin/backend/services/index_service"
	"ev-plugin/backend/services/navicat_service"
	"github.com/1340691923/eve-plugin-sdk-go/backend/web_engine"
)

type WebServer struct {
	ctx                 context.Context
	engine              *web_engine.WebEngine
	esBackUpController  *api.EsBackUpController
	esController        *api.EsController
	esCrudController    *api.EsCrudController
	esDocController     *api.EsDocController
	esIndexController   *api.EsIndexController
	esMappingController *api.EsMappingController
	esTaskController    *api.EsTaskController
}

// 依赖注入
func NewWebServer(ctx context.Context, app *web_engine.WebEngine) *WebServer {
	baseController := api.NewBaseController(response.NewResponse())
	esBackUpController := api.NewEsBackUpController(
		baseController, es_backup_service.NewEsBackUpService(
			cluser_settings_service.NewClusterSettingsService()),
	)
	esController := api.NewEsController(baseController, cat_service.NewCatService(), es_service.NewEsService())
	esCurdController := api.NewEsCrudController(baseController, navicat_service.NewNavicatService())
	esDocController := api.NewEsDocController(baseController, es_doc_service.NewEsDocService())
	esIndexController := api.NewEsIndexController(baseController, index_service.NewIndexService())
	esMapController := api.NewEsMappingController(baseController, index_service.NewIndexService())
	esTaskController := api.NewEsTaskController(baseController, es_task_service.NewEsTaskService())

	return &WebServer{
		ctx:                 ctx,
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
func NewRouter(ctx context.Context) *web_engine.WebEngine {

	//后端api
	webSvr := NewWebServer(ctx, web_engine.NewWebEngine())

	webSvr.runDslHistory()
	webSvr.runEs()
	webSvr.runEsIndex()
	webSvr.runEsTask()
	webSvr.runEsBackUp()
	webSvr.runEsCrud()
	webSvr.runEsDoc()
	webSvr.runEsMap()

	return webSvr.engine
}
