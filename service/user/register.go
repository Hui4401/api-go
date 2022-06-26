package user

import (
    sqlModel "api-go/storage/mysql/model"
    "api-go/model"
    "api-go/util/errors"
)

func Register(req *model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
    // 表单验证
    if err := registerValid(req); err != nil {
        return nil, err
    }

    user := sqlModel.User{
        Username: req.Username,
        Nickname: req.Username,
    }
    // 加密密码
    if err := user.SetPassword(req.Password); err != nil {
        return nil, err
    }
    // 创建用户
    ud := sqlModel.NewUserDao()
    if err := ud.CreateUser(&user); err != nil {
        return nil, err
    }

    return &model.UserRegisterResponse{
        User: &model.UserInfo{
            Id:        user.ID,
            Username:  user.Nickname,
            Nickname:  user.Nickname,
            CreatedAt: user.CreatedAt.Unix(),
        },
    }, nil
}

func registerValid(req *model.UserRegisterRequest) error {
    // 两次输入密码不一致
    if req.PasswordConfirm != req.Password {
        return errors.NewCodeError(errors.CodePasswordConfirmError)
    }
    // 用户名已存在
    ud := sqlModel.NewUserDao()
    user, err := ud.GetUserByUsername(req.Username)
    if err != nil {
        return err
    }
    if user != nil {
        return errors.NewCodeError(errors.CodeUserExist)
    }

    return nil
}
