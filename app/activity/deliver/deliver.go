package activity

import (
	"admin/app/activity/entity"
	"admin/app/activity/repo"
	"admin/app/activity/ucase"
	"admin/common/resp"
	"admin/init/pgsql"
	"admin/init/redis"
	"log"

	"github.com/gin-gonic/gin"
)

var pgServer = repo.NewPgRepo(pgsql.DBConn)
var redisServer = repo.NewRedis(redis.RedisPool)
var ucaseServer = ucase.NewUcase(pgServer, redisServer)

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
