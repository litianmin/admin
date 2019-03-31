package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	loginServer "admin/app/login/deliver"
	"admin/common/upload"
	"admin/middleware"
)

func mytest(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	fileC, err := file.Open()
	if err != nil {
		fmt.Println("something happened here!")
	}

	str := make([]byte, 2<<20)

	fileC.Read(str)
	newFile, _ := os.Create("./mytest.png")

	defer newFile.Close()

	newFile.Write(str)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	c.JSON(200, "不错，你已经上传成功了")

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func main() {
	r := gin.Default()

	r.POST("/login", loginServer.Login)

	r.POST("/upload", middleware.TokenIsValid, upload.GameLogoUpload)

	r.Run(":9999")
}
