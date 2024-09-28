// 自定义响应 辅助方法层
package response

import (
	"ev-tools/backend/my_error"
	"fmt"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"` //消息码
	Msg  string      `json:"msg"`  //消息提示
	Data interface{} `json:"data"` //附加信息
}

// 自定义响应方法
type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

const (
	SUCCESS = 0
	ERROR   = 500
)

const (
	SearchSuccess       = "查询成功"
	DeleteSuccess       = "删除成功"
	OperateSuccess      = "操作成功"
	LogoutSuccess       = "注销成功"
	ChangeLayoutSuccess = "修改布局成功"
)

// 正确信息
func (this *Response) Success(ctx *gin.Context, msg string, data interface{}) error {
	responseData := new(ResponseData)
	responseData.Msg = msg
	responseData.Data = data
	responseData.send(ctx, SUCCESS)
	return nil
}

// 错误信息
func (this *Response) Error(ctx *gin.Context, err error) error {

	myErr := ErrorToErrorCode(err)

	var b []byte
	b, _ = ctx.GetRawData()

	logger.DefaultLogger.Error("错误日志",
		"请求接口地址", ctx.Request.URL.Path,
		"请求body", string(b),
		"异常堆栈", fmt.Sprintf("%+v", err),
	)

	responseData := new(ResponseData)
	responseData.Msg = myErr.Error()
	responseData.send(ctx, myErr.Code())
	return nil
}

// 输出
func (this *ResponseData) send(ctx *gin.Context, code int) error {
	this.Code = code
	var err error
	if this.Code != 0 {
		ctx.JSON(http.StatusAccepted, this)
	} else {
		ctx.JSON(http.StatusOK, this)
	}

	if err != nil {
		ctx.JSON(http.StatusAccepted, map[string]interface{}{"msg": err, "code": 500})
	}
	return nil
}

// 输出
func (this *Response) Output(write io.Writer, data map[string]interface{}) error {
	b, _ := json.Marshal(data)
	write.Write(b)
	return nil
}

func ErrorToErrorCode(err error) *my_error.MyError {
	if err == nil {
		return nil
	}

	errorCode, ok := err.(*my_error.MyError)

	if ok {
		return errorCode
	}
	return my_error.NewError(err.Error(), ERROR).(*my_error.MyError)
}
