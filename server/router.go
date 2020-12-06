package server

import (
	"api-go/api"
	v1 "api-go/api/v1"
	"api-go/middleware/auth"

	"github.com/gin-gonic/gin"
)

// 路由配置
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 主页
	r.GET("/", api.Index)

	v1Group := r.Group("/api/v1")
	{
		// 用户注册
		v1Group.POST("/user/register", v1.UserRegister)

		// 用户登录
		v1Group.POST("/user/login", v1.UserLogin)

		// 需要权限
		jwtGroup := v1Group.Group("")
		jwtGroup.Use(auth.JwtRequired())
		{
			// 查看个人信息
			jwtGroup.GET("/user/me", v1.UserMe)

			// 退出登录
			jwtGroup.POST("/user/logout", v1.UserLogout)
		}
	}
	return r
}
