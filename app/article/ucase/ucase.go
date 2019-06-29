package ucase

import "admin/app/article/entity"

// MysqlServer 定义repo 那边需要实现的接口
type MysqlServer interface {
	NewArticle(data *entity.NewArticle) (IsSuccess bool, Article *entity.ArticleBaseInfo)
}

// MongoServer 定义mongoRepo
type MongoServer interface {
	NewArticle(article *entity.ArticleBaseInfo) bool
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

// NewArticle 新增文章
func (u *Ucase) NewArticle(data *entity.NewArticle) bool {
	stepOne, article := u.MysqlRepo.NewArticle(data)
	if stepOne == false {
		return false
	}

	stepTwo := u.MongoRepo.NewArticle(article)

	if stepTwo == false {
		return false
	}

	return true
}
