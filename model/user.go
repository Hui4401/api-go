package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
}

const (
	// 密码加密级别
	passwordCost = bcrypt.DefaultCost
)

// 用ID获取用户
func GetUser(ID uint) (*User, error) {
	var user User
	result := DB.First(&user, ID)
	return &user, result.Error
}

// 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
