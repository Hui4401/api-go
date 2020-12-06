package api

import (
	"github.com/gin-gonic/gin"
)

// 主页
func Index(c *gin.Context) {
	c.String(200, "================   Welcome to api-go Restful API Index Page!    https://github.com/Hui4401/api-go   ================")
}
