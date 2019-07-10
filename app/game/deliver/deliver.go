package deliver

import (
	"admin/app/game/entity"
	"admin/app/game/repo"
	"admin/app/game/ucase"
	"admin/app/game/validate"
	"admin/common/resp"
	"admin/init/pgsql"

	"github.com/gin-gonic/gin"
)

var pgServer = repo.NewPgRepo(pgsql.DBConn)
var ucaseServer = ucase.NewUcase(pgServer)

// CreateGame 创建新的游戏
func CreateGame(c *gin.Context) {
	var newGame entity.NewGame
	c.Bind(&newGame)
	isValid := validate.CreateGame(&newGame)
	if isValid == false {
		c.JSON(200, resp.ParamsErr)
	}

	isSuccess := ucaseServer.CreateGame(&newGame)
	if isSuccess {
		c.JSON(200, gin.H{
			"code": 20000,
			"msg":  "添加成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "发生未知错误",
		})
	}

}
