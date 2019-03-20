package ucase

import (
	"admin/app/login/entity"
	"fmt"
)

// RepoServer 定义repo 那边需要实现的接口
type RepoServer interface {
	LoginAuth(body *entity.LoginAuth)
}

// Ucase 定义结构体
type Ucase struct {
	repo RepoServer
}

// NewUcase 初始化 Ucase
func NewUcase(repo RepoServer) *Ucase {
	return &Ucase{repo}
}

// LoginAuth 定义 LoginAuth 方法
func (u *Ucase) LoginAuth(body *entity.LoginAuth) {
	fmt.Println(body.UserName)
}
