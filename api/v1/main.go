package v1

import (
    "api-go/model"
    "github.com/gin-gonic/gin"
    "fmt"
)

// 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
    if userID, ok := c.Get("user_id"); ok {
        if user, err := model.GetUser(*userID.(*uint)); err == nil {
            return user
        } else {
            fmt.Println(err.Error())
        }
    }
    return nil
}
