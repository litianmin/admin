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

// Login serve for login!
func Login(c *gin.Context) {
	body := entity.LoginAuth{}
	c.BindJSON(&body)
	ucaseServer.LoginAuth(&body)
	c.JSON(200, "ceshi")
}
