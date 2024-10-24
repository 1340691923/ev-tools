package api

import (
	"ev-tools/backend/response"
	"ev-tools/backend/services/alias_service"
	"ev-tools/backend/services/index_service"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto/common"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Es 索引控制器
type EsIndexController struct {
	*BaseController

	indexService *index_service.IndexService
}

func NewEsIndexController(baseController *BaseController, indexService *index_service.IndexService) *EsIndexController {
	return &EsIndexController{BaseController: baseController, indexService: indexService}
}

// @Summary 创建索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsIndexInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/CreateAction [post]
func (this *EsIndexController) CreateAction(ctx *gin.Context) {
	esIndexInfo := new(dto.EsIndexInfo)
	err := ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = this.indexService.EsIndexCreate(ctx, esI, esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 删除索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsIndexInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/DeleteAction [post]
func (this *EsIndexController) DeleteAction(ctx *gin.Context) {
	esIndexInfo := new(dto.EsIndexInfo)
	err := ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = this.indexService.EsIndexDelete(ctx, esI, esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return

}

// @Summary 获取索引配置信息
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsIndexInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/GetSettingsAction [post]
func (this *EsIndexController) GetSettingsAction(ctx *gin.Context) {
	esIndexInfo := new(dto.EsIndexInfo)
	err := ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexGetSettings(ctx, esI, esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)

}

// @Summary 获取所有的索引配置信息
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsIndexInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/GetSettingsInfoAction [post]
func (this *EsIndexController) GetSettingsInfoAction(ctx *gin.Context) {
	esIndexInfo := new(dto.EsIndexInfo)
	err := ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexGetSettingsInfo(ctx, esI, esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 获取别名
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsAliasInfo false "查询参数"
// @Success 0 {object} vo.AliasInfo
// @Router /api/es_index/GetAliasAction [post]
func (this *EsIndexController) GetAliasAction(ctx *gin.Context) {
	esAliasInfo := new(dto.EsAliasInfo)
	err := ctx.Bind(&esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(
		esAliasInfo.EsConnect,

		cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := alias_service.NewAliasService().EsIndexGetAlias(ctx, esI, esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)

}

// @Summary 移动别名到索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsAliasInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/MoveAliasToIndex [post]
func (this *EsIndexController) MoveAliasToIndex(ctx *gin.Context) {
	esAliasInfo := new(dto.EsAliasInfo)
	err := ctx.Bind(&esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esAliasInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = alias_service.NewAliasService().MoveAliasToIndex(ctx, esI, esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 新增别名到索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsAliasInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/AddAliasToIndex [post]
func (this *EsIndexController) AddAliasToIndex(ctx *gin.Context) {
	esAliasInfo := new(dto.EsAliasInfo)
	err := ctx.Bind(&esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esAliasInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = alias_service.NewAliasService().AddAliasToIndex(ctx, esI, esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 批量新增别名到索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsAliasInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/BatchAddAliasToIndex [post]
func (this *EsIndexController) BatchAddAliasToIndex(ctx *gin.Context) {
	esAliasInfo := new(dto.EsAliasInfo)
	err := ctx.Bind(&esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esAliasInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = alias_service.NewAliasService().BatchAddAliasToIndex(ctx, esI, esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 移除别名
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsAliasInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/RemoveAlias [post]
func (this *EsIndexController) RemoveAlias(ctx *gin.Context) {
	esAliasInfo := new(dto.EsAliasInfo)
	err := ctx.Bind(&esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esAliasInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = alias_service.NewAliasService().RemoveAlias(ctx, esI, esAliasInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 重建索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsReIndexInfo false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/ReindexAction [post]
func (this *EsIndexController) ReindexAction(ctx *gin.Context) {
	esReIndexInfo := new(dto.EsReIndexInfo)
	err := ctx.Bind(&esReIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esReIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexReindex(ctx, esI, esReIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 得到所有的索引名
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsConnectID false "查询参数"
// @Success 0 {object} []string
// @Router /api/es_index/IndexNamesAction [post]
func (this *EsIndexController) IndexNamesAction(ctx *gin.Context) {
	esConnectID := new(common.EsConnectID)
	err := ctx.Bind(&esConnectID)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esConnectID.EsConnectID, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexNames(ctx, esI)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 得到所有的索引数量
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsConnectID false "查询参数"
// @Success 0 {object} int
// @Router /api/es_index/IndexsCountAction [post]
func (this *EsIndexController) IndexsCountAction(ctx *gin.Context) {
	esConnectID := new(common.EsConnectID)
	err := ctx.Bind(&esConnectID)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esConnectID.EsConnectID, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexCount(ctx, esI)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 获取索引的Stats
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsIndexInfo false "查询参数"
// @Success 0 {object} []vo.Status
// @Router /api/es_index/StatsAction [post]
func (this *EsIndexController) StatsAction(ctx *gin.Context) {
	esIndexInfo := new(dto.EsIndexInfo)
	err := ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esIndexInfo.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.indexService.EsIndexStats(ctx, esI, esIndexInfo.IndexName)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 刷新索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Refresh [post]
func (this *EsIndexController) Refresh(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.CacheClear(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 清理索引缓存
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/CacheClear [post]
func (this *EsIndexController) CacheClear(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.CacheClear(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 将所有索引刷新到磁盘
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Flush [post]
func (this *EsIndexController) Flush(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.Flush(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 强制合并索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Forcemerge [post]
func (this *EsIndexController) Forcemerge(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.IndicesForcemerge(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 开启索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Open [post]
func (this *EsIndexController) Open(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.Open(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 关闭索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Close [post]
func (this *EsIndexController) Close(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.Close(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 清空索引
// @Tags es索引
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsOptimize false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_index/Empty [post]
func (this *EsIndexController) Empty(ctx *gin.Context) {
	esOptimize := new(dto.EsOptimize)
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esOptimize.EsConnect, cast.ToInt(ctx.GetHeader(util.EvUserID)))

	indexSvr := index_service.NewIndexService()
	if esOptimize.IndexName == "" {
		esOptimize.IndexName = "*"
	}
	err = indexSvr.Empty(ctx, esI, []string{esOptimize.IndexName})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return
}
