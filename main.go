package main

import (
	"github.com/gin-gonic/gin"

	activityServer "admin/app/activity/deliver"
	articleServer "admin/app/article/deliver"
	gameServer "admin/app/game/deliver"
	userServer "admin/app/user/deliver"
	"admin/common/upload"
	"admin/middleware"
)

func main() {
	r := gin.New()

	r.POST("/login", userServer.Login)

	// 图片上传
	r.POST("/upload", middleware.TokenIsValid, upload.ImgUpload)

	// 新增游戏
	r.POST("/gamecreate", middleware.TokenIsValid, gameServer.CreateGame)

	// 创建活动
	r.POST("/activity/create", middleware.TokenIsValid, activityServer.NewActivity)

	// 创建文章
	r.POST("/article/create", middleware.TokenIsValid, articleServer.NewArticle)

	r.Run(":9999")
}
