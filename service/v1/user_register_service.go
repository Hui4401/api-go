package v1

import (
    "api-go/model"
    "api-go/serializer"
)

// 用户注册所需信息
type UserRegisterService struct {
    Username        string `form:"username" binding:"required,min=3,max=10"`
    Password        string `form:"password" binding:"required,min=6,max=18"`
    PasswordConfirm string `form:"password_confirm" binding:"required,min=6,max=18"`
}

// 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
    // 两次输入密码不一致
    if service.PasswordConfirm != service.Password {
        return serializer.ErrorResponse(serializer.CodePasswordConfirmError)
    }

    // 用户名已存在
    res := model.DB.Where("username = ?", service.Username).First(&model.User{})
    if res.RowsAffected > 0 {
        return serializer.ErrorResponse(serializer.CodeUserExistError)
    }

    return nil
}

// 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
    // 表单验证
    if err := service.Valid(); err != nil {
        return err
    }

    user := model.User {
        Username: service.Username,
        Nickname: service.Username,
    }

    // 加密密码
    if err := user.SetPassword(service.Password); err != nil {
        return serializer.ErrorResponse(serializer.CodeUnknownError)
    }

    // 创建用户
    if err := model.DB.Create(&user).Error; err != nil {
        return serializer.ErrorResponse(serializer.CodeUnknownError)
    }

    // 响应信息
    data := serializer.BuildUserResponse(&user)
    return serializer.OkResponse(data)
}
