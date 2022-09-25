package main

import (
    "os"

    "github.com/gin-gonic/gin"

    "api-go/conf"
    "api-go/router"
    "api-go/storage/mysql"
    sqlModel "api-go/storage/mysql/model"
    "api-go/storage/redis"
    "api-go/util/logs"
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
        logs.PanicKvs("run server error", err)
    }
}
