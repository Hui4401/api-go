package model

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	Username  string
	Password  string
	Nickname  string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = bcrypt.DefaultCost
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
