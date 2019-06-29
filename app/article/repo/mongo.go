package repo

import (
	"admin/app/article/entity"
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

// NewArticle 新增文章
func (mg *MongoRepo) NewArticle(article *entity.ArticleBaseInfo) bool {

	// 设置删除的时间，默认设置文章展示最后的时间
	expiredTimeObj := utils.UnixFormatToTimeObj(article.EndTime)

	stmt := mg.Conn.Collection("article")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 插入的数据
	insertData := bson.M{
		"articleID":      article.ArticleID,
		"title":          article.Title,
		"type":           article.Type,
		"beginTIme":      article.BeginTime,
		"endTime":        article.EndTime,
		"displayImg":     article.DisplayImg,
		"expiredTimeObj": expiredTimeObj,
	}

	_, err := stmt.InsertOne(ctx, insertData)
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}
