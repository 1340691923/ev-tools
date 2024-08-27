package api

import (
	"ev-tools/backend/response"
	"ev-tools/backend/services/es_task_service"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Es 任务控制器
type EsTaskController struct {
	*BaseController

	taskService *es_task_service.EsTaskService
}

func NewEsTaskController(baseController *BaseController, taskService *es_task_service.EsTaskService) *EsTaskController {
	return &EsTaskController{BaseController: baseController, taskService: taskService}
}

// @Summary Es任务列表
// @Tags es任务
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.TaskList false "查询参数"
// @Success 0 {object} map[string]vo.TaskInfo
// @Router /api/es_task/ListAction [post]
func (this *EsTaskController) ListAction(ctx *gin.Context) {
	taskListReq := new(dto.TaskList)
	err := ctx.Bind(&taskListReq)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(taskListReq.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := this.taskService.TaskList(ctx, esI)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res)
}

// @Summary 取消ES任务
// @Tags es任务
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.CancelTask false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es_task/CancelAction [post]
func (this *EsTaskController) CancelAction(ctx *gin.Context) {
	cancelTask := new(dto.CancelTask)
	err := ctx.Bind(&cancelTask)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(cancelTask.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = this.taskService.Cancel(ctx, esI, cancelTask)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return

}
