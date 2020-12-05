package serializer

import "api-go/model"

// 单个用户信息
type UserData struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化单个用户信息
func BuildUserData(user *model.User) *UserData {
	return &UserData{
		Id:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// 带标签的单个用户响应信息
type UserResponse struct {
	Data *UserData `json:"user"`
}

// 序列化单个用户响应信息
func BuildUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Data: BuildUserData(user),
	}
}
