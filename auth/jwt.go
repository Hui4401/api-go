package auth

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

const (
    JwtSecretKey = "a random key"
    JwtExpiresTime = time.Hour * 24
)

// Jwt 需要编码的结构体
type Jwt struct {
    jwt.StandardClaims
    UserID uint
}
