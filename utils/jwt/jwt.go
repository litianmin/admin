package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	privateKey      string = "icomingfortianmin"
	expiredDuration int64  = 7200
	alg             string = "HS256"
	iss             string = "icoming.top"
	audUser         string = "user"
	audAmin         string = "admin"
)

type jwtTokenHeader struct {
	Alg string
	Tp  string
}

type jwtPayLoad struct {
	Iss string // 网站
	Aud string // 用户类型
	Sub uint64 // 用户id
	Exp int64  // 过期时间
	Iat int64  // 签发时间
}

// CreateJwtToken 创建jwtToken
func CreateJwtToken(userID uint64, userTp string) (string, bool) {
	var aud string

	if userID < 0 {
		return "", false
	}

	// jwt头部
	header := &jwtTokenHeader{
		Alg: alg,
		Tp:  "jwt",
	}

	switch userTp {
	case "user":
		aud = audUser
	case "admin":
		aud = audAmin
	default:
		break
	}

	thisMoment := time.Now().Unix()

	// 载荷
	payLoad := &jwtPayLoad{
		Iss: iss,
		Aud: aud,
		Sub: userID,
		Exp: thisMoment + expiredDuration,
		Iat: thisMoment,
	}

	headJSON, _ := json.Marshal(header)
	payLoadJSON, _ := json.Marshal(payLoad)

	headBase64 := base64.StdEncoding.EncodeToString(headJSON)
	payLoadBase64 := base64.StdEncoding.EncodeToString(payLoadJSON)

	mergeStr := headBase64 + "." + payLoadBase64

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(mergeStr))
	signedByte := h.Sum(nil)
	signedStr := hex.EncodeToString(signedByte)
	res := headBase64 + "." + payLoadBase64 + "." + signedStr
	return res, true
}

// ParseJwtToken 解析token,验证是否已经被修改了，是否已经过期了
// code	  0 => 没有问题，1 => 非法, 2 => 过期了
func ParseJwtToken(token *string) (code int, usrID uint64) {
	strArr := strings.Split(*token, ".")
	if len(strArr) != 3 {
		return 1, 0
	}

	mergeStr := strArr[0] + "." + strArr[1]

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(mergeStr))

	signedByte := h.Sum(nil)
	signedStr := hex.EncodeToString(signedByte)

	// 判断签名是否相同， 如果相同是合法的，如果不相同那么就是不合法的
	if signedStr != strArr[2] {
		return 1, 0
	}

	// 判断该token是否已经过期了
	payloadBytes, _ := base64.StdEncoding.DecodeString(strArr[1])
	var payload jwtPayLoad
	json.Unmarshal(payloadBytes, &payload)
	thisMoment := time.Now().Unix()
	if payload.Exp < thisMoment {
		return 2, 0
	}

	return 0, payload.Sub
}

func main() {
	str, _ := CreateJwtToken(568, "user")
	fmt.Println(str)
}
