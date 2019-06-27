package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// MongoDB mongoDB初始化
	MongoDB *mongo.Database
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Panicln(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Panicln("链接失败")
	} else {
		fmt.Println("Mongo 链接成功！")
	}
	MongoDB = client.Database("ic")
	// collection := client.Database("ic").Collection("lbs")
}

// bson.A  一般用于数组， 例如： bson.A[11, 22]
// bson.D 就是顺序很重的时候使用， 例如： bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
// bson.E 使用再bson.D里面，暂时还没用过
// bson.M 和bson.D 相反，顺序不重要， 例如: bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
