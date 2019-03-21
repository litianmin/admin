package deliver

import (
	"admin/app/login/entity"
	"admin/app/login/repo"
	"admin/app/login/ucase"
	"admin/init/mysql"

	"github.com/gin-gonic/gin"
)

var repoServer = repo.NewRepo(mysql.DBConn)
var ucaseServer = ucase.NewUcase(repoServer)

type mytest struct {
	code uint
	msg  string
}

// Login serve for login!
func Login(c *gin.Context) {
	body := entity.LoginAuth{}
	c.BindJSON(&body)
	isPass, token := ucaseServer.LoginAuth(&body)

	if isPass == false {
		c.JSON(200, gin.H{
			"ceshi": "账号密码错误",
		})
	} else {
		c.JSON(200, gin.H{
			"ceshi": token,
		})
	}
}
