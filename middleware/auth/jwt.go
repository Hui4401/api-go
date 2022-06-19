package auth

import (
    "api-go/cache"
    "api-go/serializer"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

const (
    // jwt加密秘钥
    JwtSecretKey = "a random key"
    // jwt过期时间
    JwtExpiresTime = time.Hour * 24
)

// jwt编码的结构体
type JwtClaim struct {
    jwt.StandardClaims
    UserID uint
}

// jwt中间件
func JwtRequired() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 从请求头获得token
        userToken := ctx.Request.Header.Get("token")
        // 判断请求头中是否有token
        if userToken == "" {
            ctx.JSON(200, serializer.ErrorResponse(serializer.CodeTokenNotFoundError))
            ctx.Abort()
            return
        }

        // 解码token值
        token, err := jwt.ParseWithClaims(userToken, &JwtClaim{},
            func(token *jwt.Token) (interface{}, error) {
                return []byte(JwtSecretKey), nil
            })
        if err != nil || token.Valid != true {
            // 过期或者非正确处理
            ctx.JSON(200, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
            ctx.Abort()
            return
        }

        // 判断令牌是否在黑名单里面
        if result, _ := cache.Redis.SIsMember(ctx, "jwt:baned", token.Raw).Result(); result {
            ctx.JSON(200, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
            ctx.Abort()
            return
        }

        // 用户id存入上下文
        if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
            ctx.Set("user_id", &jwtStruct.UserID)
        }

        // 将Token放入Context, 用于退出登录添加黑名单
        ctx.Set("token", token.Raw)
    }
}
