package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"api-go/constdef"
	redisModel "api-go/storage/redis/model"
	"api-go/util"
)

const (
	JwtSecretKey   = "a random key"
	JwtExpiredTime = time.Hour * 24 * 7
)

type JwtClaim struct {
	jwt.StandardClaims
	UserID uint
}

// JwtAuthRequired 通过jwt秘钥来验证用户身份
func JwtAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头获得token
		userToken := ctx.Request.Header.Get("Authorization")

		// 判断请求头中是否有token
		if userToken == "" {
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenNotFound))
			ctx.Abort()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecretKey), nil
		})
		if err != nil || token.Valid != true {
			// 过期或者token不正确
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenExpired))
			ctx.Abort()
			return
		}

		// 判断token是否已退出登录
		jd := redisModel.NewJwtDao()
		if jd.IsBanedToken(ctx, token.Raw) {
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenExpired))
			ctx.Abort()
			return
		}

		// context保存token信息
		if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
			ctx.Set("user_id", &jwtStruct.UserID)
		}
		ctx.Set("token", token.Raw)
	}
}
