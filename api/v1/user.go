package v1

import (
    "api-go/serializer"
    "api-go/service/v1"
    "github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(c *gin.Context) {
    var service v1.UserRegisterService
    var res *serializer.Response
    if err := c.ShouldBind(&service); err != nil {
        res = serializer.ErrorResponse(serializer.CodeParamError)
    } else {
        res = service.Register()
    }
    c.JSON(200, res)
}

// 用户登录
func UserLogin(c *gin.Context) {
    var service v1.UserLoginService
    var res *serializer.Response
    if err := c.ShouldBind(&service); err != nil {
        res = serializer.ErrorResponse(serializer.CodeParamError)
    } else {
        res = service.Login()
    }
    c.JSON(200, res)
}

// 查看个人信息
func UserMe(c *gin.Context) {
    var res *serializer.Response
    if user := CurrentUser(c); user != nil {
        res = serializer.OkResponse(serializer.BuildUserResponse(user))
    } else {
        res = serializer.ErrorResponse(serializer.CodeUnknownError)
    }
    c.JSON(200, res)
}

//// ChangePassword 修改密码
//func ChangePassword(c *gin.Context) {
//    user := api.CurrentUser(c)
//    var service v1.ChangePassword
//    if err := c.ShouldBind(&service); err == nil {
//        res := service.Change(user)
//        c.JSON(http.StatusOK, res.Result())
//    } else {
//        c.JSON(http.StatusOK, api.ErrorResponse(err).Result())
//    }
//}
//
//// Logout 用户注销
//func Logout(c *gin.Context) {
//    token, _ := c.Get("token")
//    tokenString := token.(string)
//
//    cache.RedisClient.SAdd("jwt:baned", tokenString)
//    c.JSON(http.StatusOK, serializer.Response{
//        Msg: "注销成功！",
//    })
//}
