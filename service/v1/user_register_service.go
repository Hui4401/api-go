package v1

import (
	"api-go/model"
	"api-go/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Username        string `form:"username" binding:"required,min=3,max=10"`
	Password        string `form:"password" binding:"required,min=6,max=18"`
	PasswordConfirm string `form:"password_confirm" binding:"required,min=6,max=18"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: serializer.UserPasswordError,
			Msg:  "两次输入的密码不相同",
		}
	}

	res := model.DB.Where("user_name = ?", service.Username).First(&model.User{})
	if res.RowsAffected > 0 {
		return &serializer.Response{
			Code: serializer.UserRepeatError,
			Msg:  "用户名被占用",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	user := model.User{
		Username: service.Username,
		Nickname: service.Username,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code: serializer.ServerPanicError,
			Msg:  "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Code: serializer.DatabaseWriteError,
			Msg:  "注册失败",
		}
	}

	return &serializer.Response{
		Data:      serializer.BuildUserResponse(user),
	}
}
