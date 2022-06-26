package main

import (
    "os"

    "github.com/gin-gonic/gin"

    "api-go/router"
    "api-go/conf"
    "api-go/storage/redis"
    "api-go/storage/mysql"
    "api-go/util/logs"
    sqlModel "api-go/storage/mysql/model"
)

func main() {
    conf.Init()
    mysql.InitMySQL(os.Getenv("MYSQL_URL"))
    sqlModel.AutoMigrate()
    redis.InitRedis(os.Getenv("REDIS_URL"))

    defer func() {
        logs.Sync()
    }()

    r := gin.Default()
    router.InitRouter(r)

    if err := r.Run(":8080"); err != nil {
        logs.PanicKvs("run server failed", "error", err)
    }
}
