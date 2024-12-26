package api

import (
	"ev-plugin/backend/dto"
	"ev-plugin/backend/response"
	"ev-plugin/backend/services/index_service"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
)

// Es 映射控制器
type EsMappingController struct {
	*BaseController

	indexService *index_service.IndexService
}

func NewEsMappingController(baseController *BaseController, indexService *index_service.IndexService) *EsMappingController {
	return &EsMappingController{BaseController: baseController, indexService: indexService}
}

// @Summary Es 映射列表
// @Tags es映射
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsMapGetProperties false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_map/ListAction [post]
func (this *EsMappingController) ListAction(ctx *gin.Context) {
	esMapGetProperties := new(dto.EsMapGetProperties)
	err := ctx.Bind(&esMapGetProperties)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esMapGetProperties.EsConnectID, util.GetEvUserID(ctx))

	res, err := this.indexService.EsMappingList(ctx, esI, esMapGetProperties)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	version, err := esI.EsVersion()

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res, "ver": version})
}

// @Summary 修改Es映射
// @Tags es映射
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.UpdateMapping false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_map/UpdateMappingAction [post]
func (this *EsMappingController) UpdateMappingAction(ctx *gin.Context) {
	updateMapping := new(dto.UpdateMapping)
	err := ctx.Bind(&updateMapping)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(updateMapping.EsConnect, util.GetEvUserID(ctx))

	res, err := this.indexService.UpdateMapping(ctx, esI, updateMapping)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}
