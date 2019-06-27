package repo

import (
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
func (mg *MongoRepo) NewOfficialActivity() {

}
