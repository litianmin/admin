package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	encryptPrivKey = "gnimoci"
	// ErrorLogPath 错误记录文件
	ErrorLogPath = "/usr/local/nginx/error"
)

// PwdSha1Encrypt 加密密码, 并且返回字符串
func PwdSha1Encrypt(pwd string) string {
	h := hmac.New(sha1.New, []byte(encryptPrivKey))
	h.Write([]byte(pwd))
	resBytes := h.Sum(nil)
	resStr := hex.EncodeToString(resBytes)
	return resStr
}

// CreatePath 创建文件路径
func CreatePath(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
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

// NowFormatToDate 返回现在时间
func NowFormatToDate() string {
	now := time.Now().Format("2006-01-02 03:04:05")
	return now
}

// NowFormatUnix 返回现在的时间戳
func NowFormatUnix() int64 {
	return time.Now().Unix()
}

// NowAfterFormatToDate 返回距离现在的时间为duration的时间格式： 例如:2006-01-02 03:04:05
func NowAfterFormatToDate(duration int64) string {
	now := time.Now().Unix() + duration
	return time.Unix(now, 0).Format("2006-01-02 15:04:05")
}

// NowFormatToYMD 返回现在时间的年月日(例如：20190101)
func NowFormatToYMD() string {
	return time.Now().Format("20060102")
}

// ErrLog 记录错误相关信息
// @logType 1=>notic, 2=>warning, 3=>fatal
func ErrLog(logType int, msg interface{}) {
	//创建文件夹
	newErrorLogPath := fmt.Sprintf("%s/%s", ErrorLogPath, NowFormatToYMD())
	os.MkdirAll(newErrorLogPath, os.ModePerm)
	fileName := fmt.Sprintf("%s/%s", newErrorLogPath, "error.log")
	//打开日志文件，不存在则创建
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	//设置输出流
	log.SetOutput(file)

	preFix := ""
	switch logType {
	case 1: // notic
		preFix = "[Notic]"
	case 2: // warning
		preFix = "[Warning]"
	case 3:
		preFix = "[Fatal]"
	}

	//日志前缀
	log.SetPrefix(preFix)

	//日志输出样式
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	str := fmt.Sprintf("%v", msg)

	log.Output(2, str)

	// 记录信息
	// log.Println(msg)
}
