package repo

import (
	"database/sql"
)

// PgRepo 定义结构体
type PgRepo struct {
	Conn *sql.DB
}

// NewPgRepo 初始化 Repo
func NewPgRepo(conn *sql.DB) *PgRepo {
	return &PgRepo{conn}
}

// LoginAuth 登陆验证
func (pg *PgRepo) LoginAuth(userName, pwd string) uint64 {
	stmt, _ := pg.Conn.Prepare("SELECT id FROM admin_user WHERE identifier = $1 AND credential = $2")
	defer stmt.Close()
	var userID uint64
	err := stmt.QueryRow(userName, pwd).Scan(&userID)
	if err != nil {
		return 0
	}
	return userID
}
