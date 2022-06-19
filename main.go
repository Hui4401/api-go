package main

import (
    "os"

    "api-go/router"
    "api-go/logs"
    "api-go/model"
    "api-go/cache"
    "api-go/conf"
)

func main() {

    conf.Init()
    logs.Init()
    model.InitMySQL(os.Getenv("MYSQL_URL"))
    cache.InitRedis(os.Getenv("REDIS_URL"))

    defer func() {
        _ = logs.Logger.Sync()
    }()

    r := router.InitRouter()

    if err := r.Run(":8000"); err != nil {
        logs.PanicKvs("run server failed", "error", err)
    }
}
