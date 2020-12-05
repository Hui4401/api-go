package v1

import (
	"api-go/middleware/auth"
	"api-go/model"
	"api-go/serializer"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 用户登录所需信息
type UserLoginService struct {
	Username string `form:"username" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required,min=6,max=18"`
}

func GenerateToken(userID uint) (string, error) {
	claim := auth.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth.JwtExpiresTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(auth.JwtSecretKey))
}

// Login 用户登录函数
func (service *UserLoginService) Login() *serializer.Response {
	var user model.User

	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeUserNotExistError)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ErrorResponse(serializer.CodePasswordError)
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeUnknownError)
	}

	data := gin.H{
		"token": token,
		"user":  serializer.BuildUserData(&user),
	}
	return serializer.OkResponse(&data)
}
