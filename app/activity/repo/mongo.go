package repo

import (
	"admin/app/activity/entity"
	"context"
	"ic/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepo 定义结构体
type MongoRepo struct {
	Conn *mongo.Database
}

// NewMongo 初始化 Repo
func NewMongo(conn *mongo.Database) *MongoRepo {
	return &MongoRepo{conn}
}

// NewOfficialActivity 创建官方活动
func (mg *MongoRepo) NewOfficialActivity(info *entity.ActivityBaseInfo) bool {
	// 活动结束后就直接删除吧
	expiredTime := info.EndTime

	// 设置删除的时间，默认设置为该活动结束的五天后
	deleteTime := expiredTime + 3600*24*2
	expiredTimeObj := utils.UnixFormatToTimeObj(deleteTime)

	stmt := mg.Conn.Collection("activity")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 插入的数据
	insertData := bson.M{
		"activityID":     info.ActivityID,
		"title":          info.Title,
		"type":           info.Type,
		"beginTIme":      info.BeginTime,
		"endTime":        info.EndTime,
		"venue":          info.Venue,
		"displayImg":     info.DisplayImg,
		"recruitNumb":    info.RecruitNumb,
		"expiredTimeObj": expiredTimeObj,
		"recruitStatus":  info.RecruitStatus,
		"locate": bson.M{
			"type":        "Point",
			"coordinates": bson.A{info.Venue.Lng, info.Venue.Lat},
		},
	}

	_, err := stmt.InsertOne(ctx, insertData)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
