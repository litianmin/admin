package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	encryptPrivKey = "gnimoci"
)

// PwdSha1Encrypt 加密密码, 并且返回字符串
func PwdSha1Encrypt(pwd string) string {
	h := hmac.New(sha1.New, []byte(encryptPrivKey))
	h.Write([]byte(pwd))
	resBytes := h.Sum(nil)
	resStr := hex.EncodeToString(resBytes)
	return resStr
}

// CreateFileAndPath 创建文件，包括文件夹, 模式都是可读的
func CreateFileAndPath(filePath, fileName string) (*os.File, bool) {

	// 首先判断是否存在文件路径
	os.MkdirAll(filePath, os.ModePerm)

	// 现在来创建文件了
	file, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return file, true
}

// CreateImgFileName 创建一个文件名称
func CreateImgFileName() string {
	rand.Seed(time.Now().Unix())
	newInt := rand.Intn(1000)

	thisMoment := time.Now().UnixNano()

	fileName := fmt.Sprintf("%d%d", thisMoment, newInt)

	return fileName
}
