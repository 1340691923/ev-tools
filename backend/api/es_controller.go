package api

import (
	"ev-tools/backend/model"
	"ev-tools/backend/response"
	"ev-tools/backend/services/cat_service"
	"ev-tools/backend/services/es_service"
	"fmt"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto/common"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/vo"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
	"github.com/cch123/elasticsql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Es 基本操作
type EsController struct {
	*BaseController

	catService *cat_service.CatService
	esService  *es_service.EsService
}

func NewEsController(baseController *BaseController, catService *cat_service.CatService, esService *es_service.EsService) *EsController {
	return &EsController{BaseController: baseController, catService: catService, esService: esService}
}

// @Summary es的cat操作
// @Tags ES
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsCat false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es/CatAction [post]
func (this *EsController) CatAction(ctx *gin.Context) {

	esCat := new(dto.EsCat)
	err := ctx.Bind(&esCat)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esCat.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	catSvr := this.catService

	var data *proto.Response

	switch esCat.Cat {
	case "CatHealth":
		data, err = catSvr.CatHealth(ctx, esI)
	case "CatShards":
		data, err = catSvr.CatShards(ctx, esI)
	case "CatCount":
		data, err = catSvr.CatCount(ctx, esI)
	case "CatAllocation":
		data, err = catSvr.CatAllocation(ctx, esI)
	case "CatAliases":
		data, err = catSvr.CatAliases(ctx, esI)
	case "CatIndices":
		data, err = catSvr.CatIndices(ctx, esI, []string{"store.size:desc"}, esCat.IndexBytesFormat)
	case "CatSegments":
		data, err = catSvr.IndicesSegmentsRequest(ctx, esI)
	case "CatStats":
		data, err = catSvr.ClusterStats(ctx, esI)
	case "Node":
		data, err = catSvr.CatNodes(ctx, esI)
	default:
		err = errors.New("未知类型")
	}

	if err != nil {
		this.Error(ctx, errors.WithStack(err))
		return
	}

	if data.StatusErr() != nil {
		this.Error(ctx, errors.WithStack(data.StatusErr()))
		return
	}

	this.Success(ctx, response.SearchSuccess, data.JsonRawMessage())
}

// @Summary SQL 转换为 DSL
// @Tags ES
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.SqlToDsl false "查询参数"
// @Success 0 {object} vo.SqlToDsl
// @Router /api/es/SqlToDslAction [post]
func (this *EsController) SqlToDslAction(ctx *gin.Context) {

	req := new(dto.SqlToDsl)
	err := ctx.Bind(&req)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	dsl, table, err := elasticsql.ConvertPretty(req.Sql)
	if err != nil {
		this.Error(ctx, errors.WithStack(err))
		return
	}
	this.Success(ctx, "转换成功!", vo.SqlToDsl{Dsl: dsl, TableName: table})
}

// @Summary 将索引恢复为可写状态   由于不可抗力，ES禁止写后，默认不会自动恢复
// @Tags ES
// @Accept application/json
// @Produce application/json
// @Param X-Token header string false "用户令牌"
// @Param object body dto.EsConnectID false "查询参数"
// @Success 0 {object} response.ResponseData
// @Router /api/es/RecoverCanWrite [post]
func (this *EsController) RecoverCanWrite(ctx *gin.Context) {
	esConnectID := new(common.EsConnectID)
	err := ctx.Bind(&esConnectID)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esI := ev_api.NewEvWrapApi(esConnectID.EsConnectID, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	err = this.esService.RecoverCanWrite(ctx, esI)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
	return

}

func (this *EsController) RunDslPOSTAction(ctx *gin.Context) {

	esRest := new(dto.EsRest)
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esI := ev_api.NewEvWrapApi(esRest.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := esI.EsRunDsl(ctx, &dto.PluginRunDsl2{
		HttpMethod: http.MethodPost,
		Path:       esRest.Path,
		Dsl:        esRest.Body,
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if res.StatusCode() != 200 && res.StatusCode() != 201 {
		this.Output(ctx.Writer, map[string]interface{}{
			"code": res.StatusCode(),
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode())),
			"data": res.JsonRawMessage(),
		})
		return
	}
	gmDslHistoryModel := model.GmDslHistoryModel{
		Uid:    cast.ToInt(ctx.GetHeader(util.EvUserID)),
		Method: http.MethodPost,
		Path:   esRest.Path,
		Body:   esRest.Body,
	}

	err = gmDslHistoryModel.Insert()

	if err != nil {
		err = errors.WithStack(err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res.JsonRawMessage())
}

func (this *EsController) RunDslGETAction(ctx *gin.Context) {

	esRest := new(dto.EsRest)
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esI := ev_api.NewEvWrapApi(esRest.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := esI.EsRunDsl(ctx, &dto.PluginRunDsl2{
		HttpMethod: http.MethodGet,
		Path:       esRest.Path,
		Dsl:        esRest.Body,
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if res.StatusCode() != 200 && res.StatusCode() != 201 {
		this.Output(ctx.Writer, map[string]interface{}{
			"code": res.StatusCode(),
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode())),
			"data": res.JsonRawMessage(),
		})
		return
	}

	gmDslHistoryModel := model.GmDslHistoryModel{
		Uid:    cast.ToInt(ctx.GetHeader(util.EvUserID)),
		Method: http.MethodGet,
		Path:   esRest.Path,
		Body:   esRest.Body,
	}

	err = gmDslHistoryModel.Insert()

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res.JsonRawMessage())
}
func (this *EsController) RunDslDELETEAction(ctx *gin.Context) {

	esRest := new(dto.EsRest)
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esI := ev_api.NewEvWrapApi(esRest.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := esI.EsRunDsl(ctx, &dto.PluginRunDsl2{
		HttpMethod: http.MethodDelete,
		Path:       esRest.Path,
		Dsl:        esRest.Body,
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}
	if res.StatusCode() != 200 && res.StatusCode() != 201 {
		this.Output(ctx.Writer, map[string]interface{}{
			"code": res.StatusCode(),
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode())),
			"data": res.JsonRawMessage(),
		})
		return
	}

	gmDslHistoryModel := model.GmDslHistoryModel{
		Uid:    cast.ToInt(ctx.GetHeader(util.EvUserID)),
		Method: http.MethodDelete,
		Path:   esRest.Path,
		Body:   esRest.Body,
	}

	err = gmDslHistoryModel.Insert()

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res.JsonRawMessage())
}
func (this *EsController) RunDslPUTAction(ctx *gin.Context) {

	esRest := new(dto.EsRest)
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esI := ev_api.NewEvWrapApi(esRest.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := esI.EsRunDsl(ctx, &dto.PluginRunDsl2{
		HttpMethod: http.MethodPut,
		Path:       esRest.Path,
		Dsl:        esRest.Body,
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}
	if res.StatusCode() != 200 && res.StatusCode() != 201 {
		this.Output(ctx.Writer, map[string]interface{}{
			"code": res.StatusCode(),
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode())),
			"data": res.JsonRawMessage(),
		})
		return
	}
	gmDslHistoryModel := model.GmDslHistoryModel{
		Uid:    cast.ToInt(ctx.GetHeader(util.EvUserID)),
		Method: http.MethodPut,
		Path:   esRest.Path,
		Body:   esRest.Body,
	}

	err = gmDslHistoryModel.Insert()

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res.JsonRawMessage())
}

func (this *EsController) RunDslHEADAction(ctx *gin.Context) {

	esRest := new(dto.EsRest)
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esI := ev_api.NewEvWrapApi(esRest.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	res, err := esI.EsRunDsl(ctx, &dto.PluginRunDsl2{
		HttpMethod: http.MethodHead,
		Path:       esRest.Path,
		Dsl:        esRest.Body,
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}
	if res.StatusCode() != 200 && res.StatusCode() != 201 {
		this.Output(ctx.Writer, map[string]interface{}{
			"code": res.StatusCode(),
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode())),
			"data": res.JsonRawMessage(),
		})
		return
	}
	gmDslHistoryModel := model.GmDslHistoryModel{
		Uid:    cast.ToInt(ctx.GetHeader(util.EvUserID)),
		Method: http.MethodHead,
		Path:   esRest.Path,
		Body:   esRest.Body,
	}

	err = gmDslHistoryModel.Insert()

	if err != nil {
		err = errors.WithStack(err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res.JsonRawMessage())
}

func (this *EsController) DslHistoryAction(ctx *gin.Context) {

	gmDslHistoryModel := model.GmDslHistoryModel{}
	err := ctx.Bind(&gmDslHistoryModel)
	if err != nil {
		this.Error(ctx, errors.WithStack(err))
		return
	}
	gmDslHistoryModel.Uid = cast.ToInt(ctx.GetHeader(util.EvUserID))

	list, err := gmDslHistoryModel.List()

	if err != nil {
		this.Error(ctx, err)
		return
	}

	count, err := gmDslHistoryModel.Count()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})

}

func (this *EsController) CleanDslHistoryAction(ctx *gin.Context) {

	gmDslHistoryModel := model.GmDslHistoryModel{}
	gmDslHistoryModel.Uid = cast.ToInt(ctx.GetHeader(util.EvUserID))
	err := gmDslHistoryModel.Clean()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, nil)
}
