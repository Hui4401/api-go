package auth

import (
	"api-go/cache"
	"api-go/serializer"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	JwtSecretKey   = "a random key"
	JwtExpiresTime = time.Hour * 24
)

// jwt编码的结构体
type JwtClaim struct {
	jwt.StandardClaims
	UserID uint
}

// jwt中间件
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		userToken := c.Request.Header.Get("token")
		// 判断请求头中是否有token
		if userToken == "" {
			c.JSON(200, serializer.ErrorResponse(serializer.CodeTokenNotFoundError))
			c.Abort()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecretKey), nil
		})
		if err != nil || token.Valid != true {
			// 过期或者非正确处理
			c.JSON(200, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
			c.Abort()
			return
		}

		// 判断令牌是否在黑名单里面
		if result, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); result {
			c.JSON(200, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
			c.Abort()
			return
		}

		// 用户id存入上下文
		if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
			c.Set("user_id", &jwtStruct.UserID)
		}

		// 将Token放入Context, 用于退出登录添加黑名单
		c.Set("token", token.Raw)
	}
}
