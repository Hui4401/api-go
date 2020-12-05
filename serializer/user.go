package serializer

import "api-go/model"

// User 用户序列化器
type User struct {
    ID        uint   `json:"id"`
    Username  string `json:"username"`
    Nickname  string `json:"nickname"`
    CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
    Data User `json:"user"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
    return User{
        ID:        user.ID,
        Username:  user.Username,
        Nickname:  user.Nickname,
        CreatedAt: user.CreatedAt.Unix(),
    }
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
    return UserResponse{
        Data: BuildUser(user),
    }
}
