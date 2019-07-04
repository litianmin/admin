package repo

import (
	"admin/app/activity/entity"
	"admin/utils"
	"context"
	"encoding/json"
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

// NewActivity 创建官方活动
// @recruitStatus 0 => 招募中， 1 => 停止招募， 招募成功/招募过期， 2 => 已删除(违规等)
func (mg *MongoRepo) NewActivity(info *entity.NewActivity) (IsSuccess bool, ActivityID string) {

	stmt := mg.Conn.Collection("activity")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 插入的数据
	insertData := bson.M{
		"title":          info.Title,
		"type":           info.Type,
		"beginTime":      info.BeginTime, // 活动开始时间
		"endTime":        info.EndTime,   // 活动结束时间， 活动结束后不能加入组队，通知活动列表也不再展示
		"venue":          info.Venue,
		"displayImg":     info.DisplayImg,
		"recruitNumb":    info.RecruitNumb,
		"hadRecruitNumb": 0,
		"recruitStatus":  0,
		"cont":           info.Cont,
		"locate": bson.M{
			"type":        "Point",
			"coordinates": bson.A{info.Venue.Lng, info.Venue.Lat},
		},
	}

	res, err := stmt.InsertOne(ctx, insertData)
	if err != nil {
		utils.ErrLog(3, err)
		return false, ""
	}

	newActivityIDByte, _ := json.Marshal(res.InsertedID) // 这个后面看看有没有其他的好办法弄出来，真他妈的恶心
	newActivityIDStr := string(newActivityIDByte)
	newActivityID := newActivityIDStr[1:(len(newActivityIDStr) - 1)]
	// fmt.Println(str)
	// fmt.Println(string(res.InsertedID))

	// myStr := "5d1da36bd95e763b0c3f792f"
	// newObj, _ := primitive.ObjectIDFromHex(myStr)
	// fmt.Println(newObj)

	return true, newActivityID
}
