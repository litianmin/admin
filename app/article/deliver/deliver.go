package deliver

import (
	"admin/app/article/entity"
	"admin/app/article/repo"
	"admin/app/article/ucase"
	"admin/common/resp"
	"admin/init/pgsql"

	"github.com/gin-gonic/gin"
)

var pgServer = repo.NewPgRepo(pgsql.DBConn)
var ucaseServer = ucase.NewUcase(pgServer)

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
