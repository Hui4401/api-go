package serializer

// 基本响应格式
type Response struct {
	Code errorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 请求成功响应
func OkResponse(data interface{}) *Response {
	return &Response{
		Code: CodeOk,
		Msg:  GetErrorMsg(CodeOk),
		Data: data,
	}
}

// 请求失败响应
func ErrorResponse(code errorCode) *Response {
	return &Response{
		Code: code,
		Msg:  GetErrorMsg(code),
		Data: nil,
	}
}
