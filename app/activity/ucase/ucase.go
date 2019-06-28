package ucase

import (
	"admin/app/activity/entity"
)

// MysqlServer 定义repo 那边需要实现的接口
type MysqlServer interface {
	NewNewOfficialActivity(data *entity.NewActivity) (IsSuccess bool, Activity *entity.ActivityBaseInfo)
}

// MongoServer 定义mongoRepo
type MongoServer interface {
	NewOfficialActivity(info *entity.ActivityBaseInfo) bool
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

// NewOfficialActivity 创建新的活动
func (u *Ucase) NewOfficialActivity(data *entity.NewActivity) bool {
	stepOne, activityBaseInfo := u.MysqlRepo.NewNewOfficialActivity(data)
	if stepOne == false {
		return false
	}

	stepTwo := u.MongoRepo.NewOfficialActivity(activityBaseInfo)
	if stepTwo == false {
		return false
	}

	return true
}
