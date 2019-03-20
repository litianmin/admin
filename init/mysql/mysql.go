package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 驱动
)

const (
	dbHost     = "127.0.0.1"
	dbUsr      = "root"
	dbPwd      = "Hbk5551412"
	dbDatabase = "icoming"
	dbPort     = "3306"
	dbCharset  = "UTF8"
)

var (
	// DBConn 初始化一个连接池
	DBConn *sql.DB
	err    error
)

func init() {
	sqlOpenstr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v", dbUsr, dbPwd, dbHost, dbPort, dbDatabase, dbCharset)
	DBConn, err = sql.Open("mysql", sqlOpenstr)
	if err != nil {
		panic(err.Error())
	}
}
