package api

import (
	"ev-plugin/backend/dto"
	"ev-plugin/backend/response"
	"ev-plugin/backend/services/es_doc_service"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
)

// ES 文档控制器
type EsDocController struct {
	*BaseController
	esDocService *es_doc_service.EsDocService
}

func NewEsDocController(baseController *BaseController, esDocService *es_doc_service.EsDocService) *EsDocController {
	return &EsDocController{BaseController: baseController, esDocService: esDocService}
}

// @Summary 删除文档数据
// @Tags es文档
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsDocDeleteRowByID false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_doc/DeleteRowByIDAction [post]
func (this *EsDocController) DeleteRowByIDAction(ctx *gin.Context) {
	esDocDeleteRowByID := new(dto.EsDocDeleteRowByID)
	err := ctx.Bind(esDocDeleteRowByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esDocDeleteRowByID.EsConnect, util.GetEvUserID(ctx))

	if len(esDocDeleteRowByID.Ids) == 0 {
		err = this.esDocService.DeleteRowByIDAction(ctx, esI, esDocDeleteRowByID)
	} else {
		err = this.esDocService.BulkDeleteByID(ctx, esI, esDocDeleteRowByID.IndexName, esDocDeleteRowByID.Type, esDocDeleteRowByID.Ids)
	}
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 修改文档
// @Tags es文档
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsDocUpdateByID false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_doc/UpdateByIDAction [post]
func (this *EsDocController) UpdateByIDAction(ctx *gin.Context) {
	esDocUpdateByID := new(dto.EsDocUpdateByID)
	err := ctx.Bind(esDocUpdateByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esDocUpdateByID.EsConnect, util.GetEvUserID(ctx))

	err = this.esDocService.EsDocUpdateByID(ctx, esI, esDocUpdateByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 新增文档
// @Tags es文档
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsDocUpdateByID false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_doc/InsertAction [post]
func (this *EsDocController) InsertAction(ctx *gin.Context) {
	esDocUpdateByID := new(dto.EsDocUpdateByID)
	err := ctx.Bind(esDocUpdateByID)

	esI := ev_api.NewEvWrapApi(esDocUpdateByID.EsConnect, util.GetEvUserID(ctx))
	if err != nil {
		this.Error(ctx, err)
		return
	}

	res, err := this.esDocService.EsDocInsert(ctx, esI, esDocUpdateByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)

}
