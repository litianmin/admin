package middleware

import (
	"admin/common/resp"
	"admin/utils/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

// TokenIsValid 判断 token 是否有效
func TokenIsValid(c *gin.Context) {

	fmt.Println("其实进到这里来了，但是没有作出任何的动作")

	token := c.Request.Header.Get("self-token")

	fmt.Println(token)

	fmt.Println("没有token")

	if len(token) == 0 {
		c.AbortWithStatusJSON(200, "something wrong")
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
}
