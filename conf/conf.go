package conf

import (
	"api-go/cache"
	"api-go/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// 初始化配置
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))

	// 启动各种连接单例
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB"))
}
