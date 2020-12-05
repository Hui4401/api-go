package serializer

import "strconv"

type errorCode int

// 通用错误码
const (
	CodeOk                 errorCode = 0
	CodeUnknownError       errorCode = -1
	CodeParamError         errorCode = 10001
	CodeTokenNotFoundError errorCode = 10002
	CodeTokenExpiredError  errorCode = 10003
)

// 用户相关的错误
const (
	// 注册错误
	CodeUserExistError       errorCode = 20001
	CodePasswordConfirmError errorCode = 20002

	// 登录错误
	CodeUserNotExistError errorCode = 21001
	CodePasswordError     errorCode = 21002
)

// 错误码与描述信息map
var msgMap = map[errorCode]string{
	CodeOk:                 "ok",
	CodeUnknownError:       "未知错误",
	CodeParamError:         "请求参数错误",
	CodeTokenNotFoundError: "需要权限",
	CodeTokenExpiredError:  "token过期或不正确",

	CodeUserExistError:       "注册失败，用户已存在",
	CodePasswordConfirmError: "注册失败，两次输入密码不一致",

	CodeUserNotExistError: "登录失败，用户名不存在",
	CodePasswordError:     "登录失败，密码错误",
}

// 根据错误码得到对应说明
func GetErrorMsg(code errorCode) string {
	msg, ok := msgMap[code]
	if !ok {
		msg = "未知错误，错误码：" + strconv.Itoa(int(code))
	}
	return msg
}
