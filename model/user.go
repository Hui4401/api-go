package model

type UserInfo struct {
    Id        uint   `json:"id"`
    Username  string `json:"username"`
    Nickname  string `json:"nickname"`
    CreatedAt int64  `json:"created_at"`
}

type UserRegisterRequest struct {
    Username        string `json:"username" binding:"required,min=3,max=20"`
    Password        string `json:"password" binding:"required,min=3,max=20"`
    PasswordConfirm string `json:"password_confirm" binding:"required,min=3,max=20"`
}

type UserRegisterResponse struct {
    User *UserInfo `json:"user"`
}

type UserLoginRequest struct {
    Username string `json:"username" binding:"required,min=3,max=20"`
    Password string `json:"password" binding:"required,min=3,max=20"`
}

type UserLoginResponse struct {
    Token string    `json:"token"`
    User  *UserInfo `json:"user"`
}
