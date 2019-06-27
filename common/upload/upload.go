package upload

import (
	"admin/common/resp"
	"admin/utils"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

const (
	imgSize2M         = 2 << 20 // 2M
	imgSize3M         = 3 << 20 // 2M
	imgSize5M         = 5 << 20
	gameLogoSrc       = "./image/game/logo/"
	gameDisplayImgSrc = "./image/game/display/"

	officialActivityImgSrc = "/image/official/acticity/"
)

// ImgUpload 图片上传处理
func ImgUpload(c *gin.Context) {

	imgTp, isExist := c.GetPostForm("imgTp")

	if isExist == false {
		c.JSON(200, resp.ParamsErr)
		return
	}

	switch imgTp {
	case "gamelogo":
		gameLogoDeal(c)
	case "gameDisplayImg":
		gameDisplayDeal(c)
	case "officialActivity":
		officialActivityImgDeal(c)
	case "officialActivityDetail":
		officialActivityDetailImgDeal(c)
	}

}

// gameLogoDeal 游戏logo上传处理
func gameLogoDeal(c *gin.Context) {

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
		fmt.Println("创建文件失败")
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

	newGameLogoSrc := gameLogoSrc[1:len(gameLogoSrc)]

	c.JSON(200, gin.H{
		"code":       20000,
		"origin_img": newGameLogoSrc + newFileName + ".png",
		"mini_img":   newGameLogoSrc + newFileName + "_mini.png",
	})

}

// gameDisplayDeal 游戏展示图片处理
func gameDisplayDeal(c *gin.Context) {
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
	newFile, isSuccess := utils.CreateFileAndPath(gameDisplayImgSrc, newFileName+".png")
	if isSuccess == false {
		return
	}
	newFile.Write(str)
	newFile.Close()

	// 生成 120*120 的缩略图
	srcImage, err := imaging.Open(gameDisplayImgSrc + newFileName + ".png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	newSrcImage := imaging.Resize(srcImage, 640, 320, imaging.Lanczos)

	imaging.Save(newSrcImage, gameDisplayImgSrc+newFileName+"_mini.png")

	newGameDisplayImgSrc := gameDisplayImgSrc[1:len(gameDisplayImgSrc)]

	c.JSON(200, gin.H{
		"code":     20000,
		"mini_img": newGameDisplayImgSrc + newFileName + "_mini.png",
	})
}

// 官方活动图片处理
func officialActivityImgDeal(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")

	// 判断文件的合法性
	isPass, newFileName := imgFileVerify(fileHeader)

	if isPass == false {
		c.JSON(200, gin.H{
			"code": 40001,
			"msg":  "文件格式不正确",
		})
		return
	}

	// 判断文件的大小
	if fileHeader.Size > imgSize5M {
		c.JSON(200, gin.H{
			"code": 40001,
			"msg":  "大小不能超过5M",
		})
		return
	}

	// TODO 删除本来的图片
	prePath := utils.NowFormatToYMD()
	fileDst := fmt.Sprintf(".%s/%s/%s", officialActivityImgSrc, prePath, newFileName)

	// 保存上传过来的图片
	err := c.SaveUploadedFile(fileHeader, fileDst)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 40001,
			"msg":  "发生未知错误",
		})
		return
	}

	// 重新设置图片的大小
	srcImage, err := imaging.Open(fileDst)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		c.JSON(200, gin.H{
			"code": 40001,
			"msg":  "发生未知错误",
		})
		return
	}
	newImage := imaging.Resize(srcImage, 640, 320, imaging.Lanczos)

	// 生成新的文件名称
	strArr := strings.Split(newFileName, ".")
	suffix := strArr[1]
	miniFileName := utils.CreateImgFileName()
	miniFileDst := fmt.Sprintf(".%s/%s/%s.%s", officialActivityImgSrc, prePath, miniFileName, suffix)

	err = imaging.Save(newImage, miniFileDst)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  miniFileDst[1:len(miniFileDst)],
	})
}

// {
// 	// errno 即错误代码，0 表示没有错误。
// 	//       如果有错误，errno != 0，可通过下文中的监听函数 fail 拿到该错误码进行自定义处理
// 	"errno": 0,

// 	// data 是一个数组，返回若干图片的线上地址
// 	"data": [
// 			"图片1地址",
// 			"图片2地址",
// 			"……"
// 	]
// }

// 官方活动细节图片处理
func officialActivityDetailImgDeal(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File

	fmt.Println(files)
	// return

	for _, file := range files {
		for _, file2 := range file {
			log.Println(file2.Filename)
		}
	}

	// 判断文件的合法性
	// isPass, newFileName := imgFileVerify(fileHeader)

	// if isPass == false {
	// 	c.JSON(200, gin.H{
	// 		"code": 40001,
	// 		"msg":  "文件格式不正确",
	// 	})
	// 	return
	// }

	// // 判断文件的大小
	// if fileHeader.Size > imgSize2M {
	// 	c.JSON(200, gin.H{
	// 		"code": 40001,
	// 		"msg":  "大小不能超过2M",
	// 	})
	// 	return
	// }

	// // TODO 删除本来的图片
	// prePath := utils.NowFormatToYMD()
	// fileDst := fmt.Sprintf(".%s/%s/%s", officialActivityImgSrc, prePath, newFileName)

	// // 保存上传过来的图片
	// err := c.SaveUploadedFile(fileHeader, fileDst)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(200, gin.H{
	// 		"code": 40001,
	// 		"msg":  "发生未知错误",
	// 	})
	// 	return
	// }

	// c.JSON(200, gin.H{
	// 	"code": 20000,
	// 	"msg":  miniFileDst[1:len(miniFileDst)],
	// })
}

// 验证文件合法性和返回新的名称
func imgFileVerify(fileInfo *multipart.FileHeader) (IsPass bool, NewFileName string) {
	strArr := strings.Split(fileInfo.Filename, ".")
	suffix := strArr[1]

	contTp := fileInfo.Header["Content-Type"][0]

	// 首先判断它的文件名后缀和Content-Type符不符合
	switch suffix {
	case "jpg":
		if contTp != "image/jpeg" {
			return false, ""
		}
	case "png":
		if contTp != "image/png" {
			return false, ""
		}
	default:
		return false, ""
	}

	newFileName := fmt.Sprintf("%s.%s", utils.CreateImgFileName(), suffix)

	return true, newFileName
}
