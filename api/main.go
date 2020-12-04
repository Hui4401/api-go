package api

import (
	"api-go/model"
	"api-go/serializer"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(200, "================   Welcome to api-go Restful API Index Page!     https://github.com/Hui4401/api-go   ================")
}


// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Code:  serializer.UserInputError,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}
