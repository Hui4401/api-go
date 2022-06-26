package model

import (
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "api-go/storage/mysql"
)

const (
    // 密码加密级别
    passwordCost = bcrypt.DefaultCost
)

type User struct {
    gorm.Model
    Username string
    Password string
    Nickname string
}

type userDao struct {
    sqlClient *gorm.DB
}

func NewUserDao() *userDao {
    return &userDao{
        sqlClient: mysql.GetClient(),
    }
}

func (d *userDao) CreateUser(u *User) error {
    if err := d.sqlClient.Create(u).Error; err != nil {
        return err
    }

    return nil
}

// GetUserByID 没有记录时返回 nil, nil
func (d *userDao) GetUserByID(ID uint) (*User, error) {
    user := &User{}
    res := d.sqlClient.First(user, ID)
    if res.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if res.Error != nil {
        return nil, res.Error
    }

    return user, nil
}

// GetUserByUsername 没有记录时返回 nil, nil
func (d *userDao) GetUserByUsername(username string) (*User, error) {
    user := &User{}
    res := d.sqlClient.Where("username = ?", username).First(user)
    if res.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if res.Error != nil {
        return nil, res.Error
    }

    return user, nil
}

func (user *User) SetPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
    if err != nil {
        return err
    }
    user.Password = string(bytes)

    return nil
}

func (user *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}
