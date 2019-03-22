package deliver

import (
	"admin/app/login/entity"
	"admin/app/login/repo"
	"admin/app/login/ucase"
	"admin/common/resp"
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
		c.JSON(200, resp.AccountPwdErr)
		c.Abort()
	}

	c.JSON(200, gin.H{
		"code":  20000,
		"token": token,
	})

}
