package main

import (
	"api-go/conf"
	"api-go/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.InitRouter()

	// 运行在8080端口
	r.Run()
}
