package server

import (
	"api-go/api"
	"api-go/api/v1"
	"api-go/middleware"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 主页.
	r.GET("/", api.Index)

	// v1 特殊情况需要 列如: 微信小程序等无法使用session维持会话的场景
	if os.Getenv("v1") == "on" {

		// 因为v2必须依赖用户模型和Redis, 所以判断是否开启了Redis和MySQL
		if os.Getenv("RIM") != "use" {
			panic(fmt.Sprintf("v1 JWT验证必须依赖于MySQL以及Redis, 请在环境变量设置RIM为'use', 并且配置MySQL和Redis的连接"))
		}


		jwtGroup := r.Group("/api/v1")
		{
			// 注册
			jwtGroup.POST("user/register", v1.UserRegister)

			// 登录
			jwtGroup.POST("user/login", v1.UserLogin)

			// 使用中间件验证.
			jwt := jwtGroup.Group("")
			jwt.Use(middleware.JwtRequired())
			{
				// 查看个人信息
				jwt.GET("user/me", v1.UserMe)
				// 修改密码
				jwt.PUT("user/changepassword", v1.ChangePassword)
				// 注销
				jwt.DELETE("user/logout", v1.Logout)
			}

		}
	}
	return r
}
