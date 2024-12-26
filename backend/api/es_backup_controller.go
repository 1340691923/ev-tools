package api

import (
	"ev-plugin/backend/dto"
	"ev-plugin/backend/response"
	"ev-plugin/backend/services/es_backup_service"
	"ev-plugin/backend/vo"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
)

// 备份控制器
type EsBackUpController struct {
	*BaseController
	esBackUpService *es_backup_service.EsBackUpService
}

func NewEsBackUpController(baseController *BaseController, esBackUpService *es_backup_service.EsBackUpService) *EsBackUpController {
	return &EsBackUpController{BaseController: baseController, esBackUpService: esBackUpService}
}

// @Summary 快照仓库列表
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsSnapshotInfo false "查询参数"
// @Success 0 {object} vo.SnapshotRepositoryList
// @Router /api/backUp/SnapshotRepositoryListAction [post]
func (this *EsBackUpController) SnapshotRepositoryListAction(ctx *gin.Context) {
	esSnapshotInfo := new(dto.EsSnapshotInfo)
	err := ctx.Bind(esSnapshotInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esSnapshotInfo.EsConnect, util.GetEvUserID(ctx))

	list, res, pathRepo, err := this.esBackUpService.SnapshotRepositoryList(ctx, esI, esSnapshotInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, vo.SnapshotRepositoryList{
		List:     list,
		Res:      res,
		PathRepo: pathRepo,
	})
}

// @Summary 新建快照仓库
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotCreateRepository false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/SnapshotCreateRepositoryAction [post]
func (this *EsBackUpController) SnapshotCreateRepositoryAction(ctx *gin.Context) {
	snapshotCreateRepository := new(dto.SnapshotCreateRepository)
	err := ctx.Bind(snapshotCreateRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotCreateRepository.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.SnapshotCreateRepository(ctx, esI, snapshotCreateRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 清理快照仓库
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CleanupeRepository false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/CleanupeRepositoryAction [post]
func (this *EsBackUpController) CleanupeRepositoryAction(ctx *gin.Context) {
	cleanupeRepository := new(dto.CleanupeRepository)
	err := ctx.Bind(cleanupeRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(cleanupeRepository.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.CleanUp(ctx, esI, cleanupeRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 删除快照仓库
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotDeleteRepository false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/SnapshotDeleteRepositoryAction [post]
func (this *EsBackUpController) SnapshotDeleteRepositoryAction(ctx *gin.Context) {
	snapshotDeleteRepository := new(dto.SnapshotDeleteRepository)
	err := ctx.Bind(snapshotDeleteRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotDeleteRepository.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.SnapshotDeleteRepository(ctx, esI, snapshotDeleteRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 创建快照
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CreateSnapshot false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/CreateSnapshotAction [post]
func (this *EsBackUpController) CreateSnapshotAction(ctx *gin.Context) {
	createSnapshot := new(dto.CreateSnapshot)
	err := ctx.Bind(createSnapshot)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(createSnapshot.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.CreateSnapshot(ctx, esI, createSnapshot)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 快照列表
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotList false "查询参数"
// @Success 0 {object} []vo.Snapshot
// @Router /api/backUp/SnapshotListAction [post]
func (this *EsBackUpController) SnapshotListAction(ctx *gin.Context) {
	snapshotList := new(dto.SnapshotList)
	err := ctx.Bind(snapshotList)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotList.EsConnect, util.GetEvUserID(ctx))

	res, err := this.esBackUpService.SnapshotList(ctx, esI, snapshotList)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 删除快照
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotDelete false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/SnapshotDeleteAction [post]
func (this *EsBackUpController) SnapshotDeleteAction(ctx *gin.Context) {
	snapshotDelete := new(dto.SnapshotDelete)
	err := ctx.Bind(snapshotDelete)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotDelete.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.SnapshotDelete(ctx, esI, snapshotDelete)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 快照详情
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotDetail false "查询参数"
// @Success 0 {object} vo.SnapshotDetail
// @Router /api/backUp/SnapshotDetailAction [post]
func (this *EsBackUpController) SnapshotDetailAction(ctx *gin.Context) {
	snapshotDetail := new(dto.SnapshotDetail)
	err := ctx.Bind(snapshotDetail)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotDetail.EsConnect, util.GetEvUserID(ctx))

	res, err := this.esBackUpService.SnapshotDetail(ctx, esI, snapshotDetail)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 将索引恢复至快照时状态
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotRestore false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/SnapshotRestoreAction [post]
func (this *EsBackUpController) SnapshotRestoreAction(ctx *gin.Context) {
	snapshotRestore := new(dto.SnapshotRestore)
	err := ctx.Bind(snapshotRestore)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotRestore.EsConnect, util.GetEvUserID(ctx))

	err = this.esBackUpService.SnapshotRestore(ctx, esI, snapshotRestore)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

// @Summary 得到快照状态
// @Tags 快照
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SnapshotStatus false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/backUp/SnapshotStatusAction [post]
func (this *EsBackUpController) SnapshotStatusAction(ctx *gin.Context) {
	snapshotStatus := new(dto.SnapshotStatus)
	err := ctx.Bind(snapshotStatus)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(snapshotStatus.EsConnect, util.GetEvUserID(ctx))

	res, err := this.esBackUpService.SnapshotStatus(ctx, esI, snapshotStatus)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}
