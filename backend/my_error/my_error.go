// 自定义异常层
package my_error

// 自定义异常结构体 实现Error方法
type MyError struct {
	code int
	msg  string
}

func NewError(text string, code int) error {
	return &MyError{code, text}
}

func (this *MyError) Error() string {
	return this.msg
}

func (this *MyError) Code() int {
	return this.code
}
