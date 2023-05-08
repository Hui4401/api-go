package user

import (
	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
	"github.com/gin-gonic/gin"

	"api-go/constdef"
	"api-go/model"
	"api-go/service/user"
	sqlModel "api-go/storage/mysql/model"
	redisModel "api-go/storage/redis/model"
)

// Register 用户注册
func Register(ctx *gin.Context) (interface{}, error) {
	req := model.UserRegisterRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(constdef.CodeParam)
	}

	res, err := user.Register(&req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Login 用户登录
func Login(ctx *gin.Context) (interface{}, error) {
	req := model.UserLoginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(constdef.CodeParam)
	}

	res, err := user.Login(&req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Me 查看个人信息
func Me(ctx *gin.Context) (interface{}, error) {
	v, ok := ctx.Get(constdef.CtxUserID)
	if !ok {
		return nil, errors.NewCodeError(constdef.CodeTokenExpired)
	}
	userID, ok := v.(uint)
	if !ok {
		logs.CtxErrorKvs(ctx, "userID covert fail, v", v)
		return nil, errors.NewCodeError(constdef.CodeUnknown)
	}

	ud := sqlModel.NewUserDao()
	u, err := ud.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.NewCodeError(constdef.CodeUserNotExist)
	}

	userInfo := &model.UserInfo{
		Id:        u.ID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		CreatedAt: u.CreatedAt.Unix(),
	}

	return userInfo, err
}

// Logout 退出登录
func Logout(ctx *gin.Context) (interface{}, error) {
	token, ok := ctx.Get(constdef.CtxUserToken)
	if !ok {
		return nil, errors.NewCodeError(constdef.CodeTokenNotFound)
	}

	jd := redisModel.NewJwtDao()
	if err := jd.BanToken(ctx, token.(string), constdef.UserTokenExpiredTime); err != nil {
		return nil, err
	}

	return nil, nil
}
