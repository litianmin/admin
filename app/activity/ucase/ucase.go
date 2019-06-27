package ucase

// MysqlServer 定义repo 那边需要实现的接口
type MysqlServer interface {
}

// MongoServer 定义mongoRepo
type MongoServer interface {
	NewOfficialActivity()
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
func (u *Ucase) NewOfficialActivity() {

}
