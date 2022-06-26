package router

import (
    "github.com/gin-gonic/gin"

    "api-go/middleware/auth"
    "api-go/handler/user"
    "api-go/handler"
    "api-go/middleware/wrapper"
)

func InitRouter(r *gin.Engine) {
    r.GET("/", handler.Index)

    userGroup := r.Group("/user")
    // 注册
    userGroup.POST("/register", wrapper.HandlerFuncWrapper(user.Register))
    // 登录
    userGroup.POST("/login", wrapper.HandlerFuncWrapper(user.Login))
    // 需要登录权限
    userAuthGroup := userGroup.Group("/", auth.JwtAuthRequired())
    {
        // 查看个人信息
        userAuthGroup.GET("/me", wrapper.HandlerFuncWrapper(user.Me))
        // 退出登录
        userAuthGroup.POST("/logout", wrapper.HandlerFuncWrapper(user.Logout))
    }
}
