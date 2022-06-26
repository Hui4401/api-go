package serializer

import "api-go/util/errors"

type Response struct {
    Code int32       `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func OkResponse(data interface{}) *Response {
    return &Response{
        Code: errors.CodeOk,
        Msg:  errors.GetErrorMsgByCode(errors.CodeOk),
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
