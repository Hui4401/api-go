package v1

import (
    "api-go/cache"
    "api-go/serializer"
    v1 "api-go/service/v1"

    "github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
    var service v1.UserRegisterService
    var res *serializer.Response
    if err := ctx.ShouldBind(&service); err != nil {
        res = serializer.ErrorResponse(serializer.CodeParamError)
    } else {
        res = service.Register()
    }
    ctx.JSON(200, res)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
    var service v1.UserLoginService
    var res *serializer.Response
    if err := ctx.ShouldBind(&service); err != nil {
        res = serializer.ErrorResponse(serializer.CodeParamError)
    } else {
        res = service.Login()
    }
    ctx.JSON(200, res)
}

// UserMe 查看个人信息
func UserMe(ctx *gin.Context) {
    var res *serializer.Response
    if user := CurrentUser(ctx); user != nil {
        res = serializer.OkResponse(serializer.BuildUserResponse(user))
    } else {
        res = serializer.ErrorResponse(serializer.CodeUnknownError)
    }
    ctx.JSON(200, res)
}

// UserLogout 退出登录
func UserLogout(ctx *gin.Context) {
    token, _ := ctx.Get("token")
    tokenString := token.(string)

    cache.Redis.SAdd(ctx, "jwt:baned", tokenString)
    ctx.JSON(200, serializer.OkResponse(nil))
}
