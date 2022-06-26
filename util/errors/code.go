package errors

type errorCode = int32

// 通用错误码
const (
    CodeOk            errorCode = 0
    CodeUnknown       errorCode = -1
    CodeParam         errorCode = 10001
    CodeTokenNotFound errorCode = 10002
    CodeTokenExpired  errorCode = 10003
)

// 用户相关的错误
const (
    CodeUserExist            errorCode = 20001
    CodePasswordConfirmError errorCode = 20002

    CodeUserNotExist  errorCode = 21001
    CodePasswordError errorCode = 21002
)

// msgMap 错误码与描述信息map
var msgMap = map[errorCode]string{
    CodeOk:            "ok",
    CodeUnknown:       "未知错误",
    CodeParam:         "请求参数错误",
    CodeTokenNotFound: "需要权限",
    CodeTokenExpired:  "token过期或不正确",

    CodeUserExist:            "注册失败，用户已存在",
    CodePasswordConfirmError: "注册失败，两次输入密码不一致",

    CodeUserNotExist:  "登录失败，用户名不存在",
    CodePasswordError: "登录失败，密码错误",
}
