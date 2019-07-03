package activity

import (
	"admin/app/activity/entity"
	"admin/app/activity/repo"
	"admin/app/activity/ucase"
	"admin/common/resp"
	"admin/init/mongo"
	"admin/init/mysql"
	"log"

	"github.com/gin-gonic/gin"
)

var mysqlServer = repo.NewMysqlRepo(mysql.DBConn)
var mongoServer = repo.NewMongo(mongo.MongoDB)
var ucaseServer = ucase.NewUcase(mysqlServer, mongoServer)

// NewActivity 创建一个新的官方活动
func NewActivity(c *gin.Context) {

	data := entity.NewActivity{}

	err := c.Bind(&data)
	if err != nil {
		log.Println(err)
		c.JSON(200, resp.ParamsErr)
		return
	}

	isSuccess := ucaseServer.NewActivity(&data)
	if isSuccess == false {
		c.JSON(200, resp.UnknownErrOccurred)
		return
	}

	c.JSON(200, resp.OperateSuccess)
}
