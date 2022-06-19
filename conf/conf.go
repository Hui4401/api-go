package conf

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func Init() {
    // 从文件读取环境变量，默认读取 .env 文件
    _ = godotenv.Load()
    gin.SetMode(os.Getenv("GIN_MODE"))
}
