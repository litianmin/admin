package deliver

import (
	"admin/app/user/entity"
	"admin/app/user/repo"
	"admin/app/user/ucase"
	"admin/app/user/validate"
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

	validate := validate.Login(&body)
	if validate == false {
		c.JSON(200, resp.ParamsErr)
		return
	}

	isPass, token := ucaseServer.LoginAuth(&body)

	if isPass == false {
		c.JSON(200, resp.AccountPwdErr)
		return
	}

	dataBack := map[string]string{
		"token": token,
	}
	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  dataBack,
	})

}
