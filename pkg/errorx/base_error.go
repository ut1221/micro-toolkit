// Package errorx
// @Author PTJ 2024-05-17 09:44:05
// @Description: 定义错误类型
package errorx

const defaultCode = 1000

type CodeError struct {
	Code int `json:"code"`

	Message string `json:"message"`

	Data interface{} `json:"data,omitempty"`
}

func NewCodeErr(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDataErr(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Message: msg, Data: data}
}

func NewDefErr(msg string) error {
	return NewCodeErr(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

// CodeErrorResponse
// @Description: 响应返回值
type CodeErrorResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`

	Data interface{} `json:"data,omitempty"`
}

// Info
//
//	@Description: 返回相关参数
//	@Author PTJ 2024-05-17 09:44:54
//	@receiver e
//	@return *CodeErrorResponse
func (e *CodeError) Info() *CodeErrorResponse {

	return &CodeErrorResponse{

		Code: e.Code,

		Message: e.Message,

		Data: e.Data,
	}

}
