package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Hui4401/gopkg/logs"

	"api-go/conf"
	"api-go/router"
)

func main() {
	conf.Init()

	defer func() {
		logs.Sync()
	}()

	r := gin.Default()
	router.InitRouter(r)

	if err := r.Run(":8080"); err != nil {
		logs.PanicKvs("run server error", err)
	}
}
