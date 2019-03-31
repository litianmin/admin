package upload

import (
	"admin/utils"
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

const (
	imgSize2M   = 2 << 20 // 2M
	gameLogoSrc = "./image/game/logo/"
)

// GameLogoUpload 游戏logo上传处理
func GameLogoUpload(c *gin.Context) {
	// single file
	fileHeader, _ := c.FormFile("file")

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("打开文件失败")
	}

	// 创建2M空间
	str := make([]byte, imgSize2M)
	file.Read(str) // 读进去

	newFileName := utils.CreateImgFileName()
	newFile, isSuccess := utils.CreateFileAndPath(gameLogoSrc, newFileName+".png")
	if isSuccess == false {
		return
	}
	newFile.Write(str)
	newFile.Close()

	// 生成 120*120 的缩略图
	srcImage, err := imaging.Open(gameLogoSrc + newFileName + ".png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	newSrcImage := imaging.Resize(srcImage, 120, 120, imaging.Lanczos)

	imaging.Save(newSrcImage, gameLogoSrc+newFileName+"_mini.png")

	c.JSON(200, gin.H{
		"code":       20000,
		"origin_img": gameLogoSrc + newFileName + ".png",
		"mini_img":   gameLogoSrc + newFileName + "_mini.png",
	})

}
