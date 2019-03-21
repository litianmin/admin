package repo

import (
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
func (r *Repo) LoginAuth(userName, pwd string) uint64 {
	stmt, _ := r.Conn.Prepare("SELECT id FROM admin_user WHERE identifier = ? AND credential = ?")
	defer stmt.Close()
	var userID uint64
	err := stmt.QueryRow(userName, pwd).Scan(&userID)
	if err != nil {
		return 0
	}
	return userID
}
