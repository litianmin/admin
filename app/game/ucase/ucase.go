package ucase

import (
	"admin/app/game/entity"
)

// PgServer pg 服务
type PgServer interface {
	CreateGame(*entity.NewGame) bool
}

// Ucase 定义结构体
type Ucase struct {
	PgRepo PgServer
}

// NewUcase 初始化 Ucase
func NewUcase(pgRepo PgServer) *Ucase {
	return &Ucase{pgRepo}
}

// CreateGame 定义 CreateGame 方法
func (u *Ucase) CreateGame(g *entity.NewGame) bool {
	return u.PgRepo.CreateGame(g)
}
