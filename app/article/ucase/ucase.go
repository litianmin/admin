package ucase

import "admin/app/article/entity"

// PgServer pg服务
type PgServer interface {
	NewArticle(data *entity.NewArticle) bool
}

// Ucase 定义结构体
type Ucase struct {
	PgRepo PgServer
}

// NewUcase 初始化 Ucase
func NewUcase(pgRepo PgServer) *Ucase {
	return &Ucase{pgRepo}
}

// NewArticle 新增文章
func (u *Ucase) NewArticle(data *entity.NewArticle) bool {
	return u.PgRepo.NewArticle(data)
}
