package router

import (
	"api-go/api"
	v1 "api-go/router/v1"

	"github.com/gin-gonic/gin"
)

// 初始化路由配置
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 主页
	r.GET("/", api.Index)

	// api服务
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", api.Index)

		// 添加v1路由
		v1.AddV1Group(apiGroup)
	}

	return r
}
