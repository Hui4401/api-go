package v1

import (
    v1 "api-go/handler/v1"
    "api-go/middleware/auth"

    "github.com/gin-gonic/gin"
)

func AddV1Group(r *gin.RouterGroup) {
    v1Group := r.Group("/v1")
    {
        // 用户注册
        v1Group.POST("/user/register", v1.UserRegister)

        // 用户登录
        v1Group.POST("/user/login", v1.UserLogin)

        // 需要权限
        jwtGroup := v1Group.Group("/")
        jwtGroup.Use(auth.JwtRequired())
        {
            // 查看个人信息
            jwtGroup.GET("/user/me", v1.UserMe)

            // 退出登录
            jwtGroup.POST("/user/logout", v1.UserLogout)
        }
    }
}
