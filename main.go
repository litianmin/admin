package main

import (
	"github.com/gin-gonic/gin"

	gameServer "admin/app/game/deliver"
	userServer "admin/app/user/deliver"
	"admin/common/upload"
	"admin/middleware"
)

func main() {
	r := gin.Default()

	r.POST("/login", userServer.Login)

	r.POST("/upload", middleware.TokenIsValid, upload.ImgUpload)

	r.POST("/gamecreate", middleware.TokenIsValid, gameServer.CreateGame)

	r.Run(":9999")
}
