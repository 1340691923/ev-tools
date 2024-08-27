package api

import (
	"ev-tools/backend/response"
	"ev-tools/backend/services/navicat_service"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

// ES CRUD操作
type EsCrudController struct {
	*BaseController

	navicatService *navicat_service.NavicatService
}

func NewEsCrudController(baseController *BaseController, navicatService *navicat_service.NavicatService) *EsCrudController {
	return &EsCrudController{BaseController: baseController, navicatService: navicatService}
}

// @Summary 可视化筛选获取数据
// @Tags 可视化筛选
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CrudFilter false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_crud/GetList [post]
func (this *EsCrudController) GetList(ctx *gin.Context) {
	crudFilter := new(dto.CrudFilter)
	err := ctx.Bind(crudFilter)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(crudFilter.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))
	if err != nil {
		this.Error(ctx, err)
		return
	}
	navicatSvr := this.navicatService

	res, count, err := navicatSvr.CrudGetList(ctx, esI, crudFilter)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res, "count": count})
}

// @Summary 可视化GetDSL
// @Tags 可视化筛选
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CrudFilter false "查询参数"
// @Success 0 {object} navicat_service.Search
// @Router /api/es_crud/GetDSL [post]
func (this *EsCrudController) GetDSL(ctx *gin.Context) {
	crudFilter := new(dto.CrudFilter)
	err := ctx.Bind(crudFilter)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	navicatSvr := this.navicatService

	res, err := navicatSvr.CrudGetDSL(ctx, crudFilter)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res})
}

// @Summary 下载navicat查询excel
// @Tags 可视化筛选
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CrudFilter false "查询参数"
// @Router /api/es_crud/Download [post]
func (this *EsCrudController) Download(ctx *gin.Context) {

	crudFilter := new(dto.CrudFilter)
	err := ctx.Bind(crudFilter)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(crudFilter.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))
	if err != nil {
		this.Error(ctx, err)
		return
	}
	navicatSvr := this.navicatService

	downloadFileName, titleList, searchData, err := navicatSvr.CrudDownload(ctx, esI, crudFilter)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	err = this.navicatService.DownloadExcel(downloadFileName, titleList, searchData, ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	return
}
