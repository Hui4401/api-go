package util

import (
	"github.com/Hui4401/gopkg/errors"

	"api-go/constdef"
)

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OkResponse(data interface{}) *Response {
	return &Response{
		Code: constdef.CodeOk,
		Msg:  errors.GetErrorMsgByCode(constdef.CodeOk),
		Data: data,
	}
}

func ErrorResponse(err error) *Response {
	return &Response{
		Code: errors.GetErrorCode(err),
		Msg:  errors.GetErrorMsg(err),
		Data: nil,
	}
}

func ErrorResponseByCode(code int32) *Response {
	return &Response{
		Code: code,
		Msg:  errors.GetErrorMsgByCode(code),
		Data: nil,
	}
}
