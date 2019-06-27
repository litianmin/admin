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

// NewOfficialActivity 创建一个新的官方活动
func NewOfficialActivity(c *gin.Context) {

	data := entity.NewActivity{}

	err := c.Bind(&data)
	if err != nil {
		log.Println(err)
		c.JSON(200, resp.ParamsErr)
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  data,
	})

	// c.JSON(200, gin.H{
	// 	"code": 20000,
	// 	"msg":  "something new",
	// })
}
