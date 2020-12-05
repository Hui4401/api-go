package api

import (
    "api-go/model"
    "github.com/gin-gonic/gin"
)

// Index 主页
func Index(c *gin.Context) {
    c.String(200, "================   Welcome to api-go Restful API Index Page!     https://github.com/Hui4401/api-go   ================")
}


// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
    if user, _ := c.Get("user"); user != nil {
        if u, ok := user.(*model.User); ok {
            return u
        }
    }
    return nil
}
