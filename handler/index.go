package handler

import (
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.String(200, "================   Welcome to api-go Index Page!    https://github.com/Hui4401/api-go   ================")
}
