package repo

import (
	"admin/app/login/entity"
	"database/sql"
)

// Repo 定义结构体
type Repo struct {
	Conn *sql.DB
}

// NewRepo 初始化 Repo
func NewRepo(conn *sql.DB) *Repo {
	return &Repo{conn}
}

// LoginAuth 在数据库进行登陆验证
func (r *Repo) LoginAuth(body *entity.LoginAuth) {

}
