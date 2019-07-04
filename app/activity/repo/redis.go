package repo

import (
	"admin/app/activity/entity"
	"admin/utils"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// RedisRepo redis库
type RedisRepo struct {
	Conn *redis.Pool
}

// NewRedis 初始化 Repo
func NewRedis(conn *redis.Pool) *RedisRepo {
	return &RedisRepo{conn}
}

// NewActivity 创建新的活动
func (rd *RedisRepo) NewActivity(activityID string, data *entity.NewActivity) bool {
	c := rd.Conn.Get()
	teamKey := fmt.Sprintf("activity:teammate:%s", activityID)

	// 创建一个空的有序集合
	_, err := c.Do("ZADD", teamKey, 0, 0)

	if err != nil {
		utils.ErrLog(3, err)
		return false
	}
	// 设置过期时间，为活动结束后两天
	deleteTime := data.EndTime + 3600*24*2
	_, err = c.Do("EXPIREAT", teamKey, deleteTime)
	if err != nil {
		utils.ErrLog(3, err)
		return false
	}

	return true
}
