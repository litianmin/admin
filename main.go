package main

import (
	"github.com/gin-gonic/gin"

	loginServer "admin/app/login/deliver"
)

func main() {
	r := gin.Default()

	r.POST("/login", loginServer.Login)

	r.Run(":9999")
}
