package errors

import "errors"

type errorWithCode struct {
    error
    code int32
}

func NewCodeError(code int32) error {
    return errorWithCode{
        code: code,
    }
}

func NewCodeErrorWithMsg(code int32, msg string) error {
    return errorWithCode{
        error: errors.New(msg),
        code:  code,
    }
}

func (err errorWithCode) Error() string {
    if err.error != nil {
        return err.error.Error()
    }

    return GetErrorMsgByCode(err.code)
}

func GetErrorCode(err error) int32 {
    if codeError, ok := err.(errorWithCode); ok {
        return codeError.code
    }
    return CodeUnknown
}

func GetErrorMsg(err error) string {
    if codeError, ok := err.(errorWithCode); ok {
        return codeError.Error()
    }
    return msgMap[CodeUnknown]
}

func GetErrorMsgByCode(code int32) string {
    msg, ok := msgMap[code]
    if !ok {
        return msgMap[CodeUnknown]
    }
    return msg
}
