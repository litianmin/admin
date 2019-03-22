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

// Login serve for login!
// 验证成功， 返回token，反则返回空
func Login(c *gin.Context) {
	body := entity.LoginAuth{}
	c.BindJSON(&body)
	isPass, token := ucaseServer.LoginAuth(&body)

	if isPass == false {
		c.JSON(200, resp.AccountPwdErr)
		c.Abort()
	}

	dataBack := map[string]string{
		"token": token,
	}
	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  dataBack,
	})

}
