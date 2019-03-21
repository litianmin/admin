package ucase

import (
	"admin/app/login/entity"
	"admin/utils"
	"admin/utils/jwt"
)

// RepoServer 定义repo 那边需要实现的接口
type RepoServer interface {
	LoginAuth(userName, pwd string) uint64
}

// Ucase 定义结构体
type Ucase struct {
	Repo RepoServer
}

// NewUcase 初始化 Ucase
func NewUcase(repo RepoServer) *Ucase {
	return &Ucase{repo}
}

// LoginAuth 定义 LoginAuth 方法
func (u *Ucase) LoginAuth(body *entity.LoginAuth) (isPass bool, jwtToken string) {
	userName := body.UserName
	pwd := body.Pwd
	pwd = utils.PwdSha1Encrypt(pwd)

	userID := u.Repo.LoginAuth(userName, pwd)
	if userID == 0 {
		return false, ""
	}
	token, _ := jwt.CreateJwtToken(userID, "admin")
	return true, token
}
