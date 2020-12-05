package v1

import (
    "api-go/api"
    "api-go/cache"
    "api-go/serializer"
    "api-go/service/v1"
    "github.com/gin-gonic/gin"
    "net/http"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
    var service v1.UserRegisterService
    if err := c.ShouldBind(&service); err == nil {
        res := service.Register()
        c.JSON(200, res.Result())
    } else {
        c.JSON(200, api.ErrorResponse(err).Result())
    }
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
    var service v1.UserLoginService
    if err := c.ShouldBind(&service); err == nil {
        res := service.Login()
        c.JSON(200, res.Result())
    } else {
        c.JSON(200, api.ErrorResponse(err).Result())
    }
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
    user := api.CurrentUser(c)
    res := serializer.Response{Data: serializer.BuildUserResponse(*user)}
    c.JSON(http.StatusOK, res.Result())
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
    user := api.CurrentUser(c)
    var service v1.ChangePassword
    if err := c.ShouldBind(&service); err == nil {
        res := service.Change(user)
        c.JSON(http.StatusOK, res.Result())
    } else {
        c.JSON(http.StatusOK, api.ErrorResponse(err).Result())
    }
}

// Logout 用户注销
func Logout(c *gin.Context) {
    token, _ := c.Get("token")
    tokenString := token.(string)

    cache.RedisClient.SAdd("jwt:baned", tokenString)
    c.JSON(http.StatusOK, serializer.Response{
        Msg: "注销成功！",
    })

}
