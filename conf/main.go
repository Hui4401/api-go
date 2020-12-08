package conf

import (
	"api-go/cache"
	"api-go/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 初始化配置
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))

	// 启动各种连接单例
	model.Database(os.Getenv("MYSQL_URL"))
	cache.Redis(os.Getenv("REDIS_URL"))
}
