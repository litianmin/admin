package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
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
