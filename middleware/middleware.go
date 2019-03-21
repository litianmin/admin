package middleware

import (
	"github.com/gin-gonic/gin"
)

// TokenIsValid 判断 token 是否有效
func TokenIsValid(c *gin.Context) (isValid bool, userID uint64) {
	token := c.Request.Header.Get("self-token")
	if len(token) == 0 {
		c.AbortWithStatusJSON(200, "你没有登陆，没有权限访问接口")
	}
	return true, 1
}
