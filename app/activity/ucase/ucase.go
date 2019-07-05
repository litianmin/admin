package ucase

import (
	"admin/app/activity/entity"
	"fmt"
)

// PgServer 初始化pgsql
type PgServer interface {
	NewActivity(data *entity.NewActivity) (IsSuccess bool, NewActivityID int64)
}

// RedisServer redis 服务集
type RedisServer interface {
	NewActivity(activityID string, data *entity.NewActivity) bool
}

// Ucase 定义结构体
type Ucase struct {
	PgRepo    PgServer
	RedisRepo RedisServer
}

// NewUcase 初始化 Ucase
func NewUcase(pgRepo PgServer, redisRepo RedisServer) *Ucase {
	return &Ucase{pgRepo, redisRepo}
}

// NewActivity 创建新的活动
func (u *Ucase) NewActivity(data *entity.NewActivity) bool {

	stepOne, newActivityID := u.PgRepo.NewActivity(data)
	if stepOne == false {
		return false
	}

	fmt.Println(newActivityID)

	// stepTwo := u.RedisRepo.NewActivity(newActivityID, data)

	// if stepTwo == false {
	// return false
	// }

	return stepOne
}
