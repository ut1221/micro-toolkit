package response

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-zero/pkg/errorx"

	"net/http"
)

type Body struct {
	Code int `json:"code"`

	Message string `json:"msg"`

	Data interface{} `json:"data,omitempty"`
}

// Response
//
//	@Description: 统一封装成功响应值
//	@Author PTJ 2024-05-17 09:45:41
//	@param w
//	@param resp
//	@param err
func Response(w http.ResponseWriter, resp interface{}, err error) {

	var body Body

	if err != nil {

		var e *errorx.CodeError
		switch {

		case errors.As(err, &e): //业务输出错误

			body.Code = e.Code

			body.Message = e.Message

			body.Data = e.Data
			//body.Data = e.Data()

		default: //系统未知错误

			body.Code = 1

			body.Message = err.Error()
		}

	} else {

		body.Code = 200

		body.Message = "请求成功!"

		body.Data = resp

	}

	httpx.OkJson(w, body)

}
