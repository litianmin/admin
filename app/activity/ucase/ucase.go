package ucase

import (
	"admin/app/activity/entity"
)

// MysqlServer 定义repo 那边需要实现的接口
type MysqlServer interface {
}

// MongoServer 定义mongoRepo
type MongoServer interface {
	NewActivity(data *entity.NewActivity) bool
}

// Ucase 定义结构体
type Ucase struct {
	MysqlRepo MysqlServer
	MongoRepo MongoServer
}

// NewUcase 初始化 Ucase
func NewUcase(mysqlRepo MysqlServer, mongoRepo MongoServer) *Ucase {
	return &Ucase{mysqlRepo, mongoRepo}
}

// NewActivity 创建新的活动
func (u *Ucase) NewActivity(data *entity.NewActivity) bool {

	isSuccess := u.MongoRepo.NewActivity(data)
	if isSuccess == false {
		return false
	}

	return true
}
