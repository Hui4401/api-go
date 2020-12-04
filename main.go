package main

import (
	"api-go/conf"
	"api-go/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	// 运行在8080端口
	r.Run()
}
