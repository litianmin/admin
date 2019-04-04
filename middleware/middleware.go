package middleware

import (
	"admin/common/resp"
	"admin/utils/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

// TokenIsValid 判断 token 是否有效
func TokenIsValid(c *gin.Context) {
	token := c.Request.Header.Get("self-token")

	if len(token) == 0 {
		c.AbortWithStatusJSON(200, resp.TokenWithout)
		return
	}

	// 判断token是否有效
	authCode, _ := jwt.ParseJwtToken(&token)
	switch authCode {
	case 0:
		return
	case 1:
		c.AbortWithStatusJSON(200, resp.TokenIllegal)
		return
	case 2:
		c.AbortWithStatusJSON(200, resp.TokenExpired)
		return
	}
	fmt.Println("yinggai bu")
}
