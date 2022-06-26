package wrapper

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "api-go/serializer"
)

type handlerFunc func(ctx *gin.Context) (interface{}, error)

func HandlerFuncWrapper(handlerFunc handlerFunc) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        data, err := handlerFunc(ctx)
        if err != nil {
            ctx.JSON(http.StatusOK, serializer.ErrorResponse(err))
            return
        }
        ctx.JSON(http.StatusOK, serializer.OkResponse(data))
        return
    }
}
