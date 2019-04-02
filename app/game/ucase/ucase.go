package ucase

import (
	"admin/app/game/entity"
)

// RepoServer 定义repo 那边需要实现的接口
type RepoServer interface {
	CreateGame(*entity.NewGame) bool
}

// Ucase 定义结构体
type Ucase struct {
	Repo RepoServer
}

// NewUcase 初始化 Ucase
func NewUcase(repo RepoServer) *Ucase {
	return &Ucase{repo}
}

// CreateGame 定义 CreateGame 方法
func (u *Ucase) CreateGame(g *entity.NewGame) bool {
	return u.Repo.CreateGame(g)
}
