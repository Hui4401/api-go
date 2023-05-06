package user

import (
	"time"

	"github.com/Hui4401/gopkg/errors"
	"github.com/dgrijalva/jwt-go"

	"api-go/constdef"
	"api-go/middleware/auth"
	"api-go/model"
	sqlModel "api-go/storage/mysql/model"
)

// Login 用户登录函数
func Login(req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	ud := sqlModel.NewUserDao()
	user, err := ud.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.NewCodeError(constdef.CodeUserNotExist)
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.NewCodeError(constdef.CodePasswordError)
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &model.UserLoginResponse{
		Token: token,
		User: &model.UserInfo{
			Id:        user.ID,
			Username:  user.Nickname,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt.Unix(),
		},
	}, nil
}

func generateToken(uid uint) (string, error) {
	claim := auth.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth.JwtExpiredTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(auth.JwtSecretKey))
}
