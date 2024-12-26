package api

import (
	"bytes"
	"ev-plugin/backend/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// 父控制器结构体
type BaseController struct {
	*response.Response
}

func NewBaseController(response *response.Response) *BaseController {
	return &BaseController{Response: response}
}

func (this *BaseController) getPostBody(ctx *gin.Context) []byte {
	body, _ := ctx.GetRawData()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body
}
