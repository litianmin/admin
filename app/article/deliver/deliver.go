package deliver

import (
	"admin/app/article/entity"
	"admin/app/article/repo"
	"admin/app/article/ucase"
	"admin/common/resp"

	"github.com/gin-gonic/gin"

	"admin/init/mongo"
	"admin/init/mysql"
)

var mysqlServer = repo.NewMysqlRepo(mysql.DBConn)
var mongoServer = repo.NewMongo(mongo.MongoDB)
var ucaseServer = ucase.NewUcase(mysqlServer, mongoServer)

// NewArticle 新增文章
func NewArticle(c *gin.Context) {
	data := entity.NewArticle{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(200, resp.ParamsErr)
		return
	}

	isSuccess := ucaseServer.NewArticle(&data)
	if isSuccess == false {
		c.JSON(200, resp.UnknownErrOccurred)
		return
	}

	c.JSON(200, resp.OperateSuccess)
}
