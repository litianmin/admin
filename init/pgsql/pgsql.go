package pgsql

import (
	"database/sql"
	"fmt"
	"log"

	// 初始化
	_ "github.com/lib/pq"
)

const (
	dbHost = "localhost"
	// dbUsr      = "postgres"
	dbUsr      = "litianmin"
	dbPwd      = "Hbk5551412"
	dbDatabase = "postgres"
)

// DBConn 初始化一个连接池
var (
	DBConn *sql.DB
	err    error
)

func init() {

	sqlOpenStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", dbUsr, dbPwd, dbDatabase)

	DBConn, err = sql.Open("postgres", sqlOpenStr)
	if err != nil {
		log.Panic(err)
	}
}
