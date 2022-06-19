package router

import (
    "api-go/handler"
    v1 "api-go/router/v1"

    "github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", handler.Index)

    apiGroup := r.Group("/handler")
    {
        apiGroup.GET("/", handler.Index)

        v1.AddV1Group(apiGroup)
    }

    return r
}
