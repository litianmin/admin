package main

import (
	"github.com/gin-gonic/gin"

	activityServer "admin/app/activity/deliver"
	gameServer "admin/app/game/deliver"
	userServer "admin/app/user/deliver"
	"admin/common/upload"
	"admin/middleware"
)

func main() {
	r := gin.New()

	r.POST("/login", userServer.Login)

	r.POST("/upload", middleware.TokenIsValid, upload.ImgUpload)

	r.POST("/gamecreate", middleware.TokenIsValid, gameServer.CreateGame)

	// activity
	r.POST("/officialActivity/create", middleware.TokenIsValid, activityServer.NewOfficialActivity)

	r.Run(":9999")
}
