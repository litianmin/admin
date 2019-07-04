package ucase

import (
	"admin/app/activity/entity"
)

// MysqlServer 定义repo 那边需要实现的接口
type MysqlServer interface {
}

// MongoServer 定义mongoRepo
type MongoServer interface {
	NewActivity(data *entity.NewActivity) (IsSuccess bool, ActivityID string)
}

// RedisServer redis 服务集
type RedisServer interface {
	NewActivity(activityID string, data *entity.NewActivity) bool
}

// Ucase 定义结构体
type Ucase struct {
	MysqlRepo MysqlServer
	MongoRepo MongoServer
	RedisRepo RedisServer
}

// NewUcase 初始化 Ucase
func NewUcase(mysqlRepo MysqlServer, mongoRepo MongoServer, redisRepo RedisServer) *Ucase {
	return &Ucase{mysqlRepo, mongoRepo, redisRepo}
}

// NewActivity 创建新的活动
func (u *Ucase) NewActivity(data *entity.NewActivity) bool {

	stepOne, newActivityID := u.MongoRepo.NewActivity(data)
	if stepOne == false {
		return false
	}

	stepTwo := u.RedisRepo.NewActivity(newActivityID, data)

	if stepTwo == false {
		return false
	}

	return true
}
