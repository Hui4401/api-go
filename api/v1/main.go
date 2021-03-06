package v1

import (
	"api-go/model"

	"github.com/gin-gonic/gin"
)

// 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if userID, ok := c.Get("user_id"); ok {
		if user, err := model.GetUser(*userID.(*uint)); err == nil {
			return user
		}
	}
	return nil
}
