package repo

import (
	"database/sql"
)

// MysqlRepo 定义结构体
type MysqlRepo struct {
	Conn *sql.DB
}

// NewMysqlRepo 初始化 Repo
func NewMysqlRepo(conn *sql.DB) *MysqlRepo {
	return &MysqlRepo{conn}
}
