package user

import (
    "github.com/gin-gonic/gin"

    "api-go/service/user"
    sqlModel "api-go/storage/mysql/model"
    "api-go/model"
    "api-go/util/errors"
    redisModel "api-go/storage/redis/model"
)

// Register 用户注册
func Register(ctx *gin.Context) (interface{}, error) {
    req := model.UserRegisterRequest{}
    if err := ctx.ShouldBind(&req); err != nil {
        return nil, errors.NewCodeError(errors.CodeParam)
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
        return nil, errors.NewCodeError(errors.CodeParam)
    }

    res, err := user.Login(&req)
    if err != nil {
        return nil, err
    }

    return res, nil
}

// Me 查看个人信息
func Me(ctx *gin.Context) (interface{}, error) {
    u, err := getCurrentUser(ctx)
    if err != nil {
        return nil, err
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
    token, ok := ctx.Get("token")
    if !ok {
        return nil, errors.NewCodeError(errors.CodeTokenNotFound)
    }

    jd := redisModel.NewJwtDao()
    if err := jd.BanToken(ctx, token.(string)); err != nil {
        return nil, err
    }

    return nil, nil
}

// getCurrentUser 获取当前用户
func getCurrentUser(ctx *gin.Context) (*sqlModel.User, error) {
    if uid, ok := ctx.Get("user_id"); ok {
        ud := sqlModel.NewUserDao()
        u, err := ud.GetUserByID(*uid.(*uint))
        if err != nil {
            return nil, err
        }
        if u == nil {
            return nil, errors.NewCodeError(errors.CodeUserNotExist)
        }
        return u, nil
    }

    return nil, errors.NewCodeError(errors.CodeTokenExpired)
}
