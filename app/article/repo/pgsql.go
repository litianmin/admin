package repo

import (
	"admin/app/article/entity"
	"admin/utils"
	"context"
	"database/sql"
	"time"
)

// PgRepo 定义结构体
type PgRepo struct {
	Conn *sql.DB
}

// NewPgRepo 初始化 Repo
func NewPgRepo(conn *sql.DB) *PgRepo {
	return &PgRepo{conn}
}

// NewArticle 新增文章
func (pg *PgRepo) NewArticle(data *entity.NewArticle) bool {
	now := utils.NowFormatUnix()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 开启事务
	tx, err := pg.Conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		utils.ErrLog(3, err)
		return false
	}

	var newArticleID uint64

	err = tx.QueryRow("INSERT INTO article(title, article_type, display_img, begin_time, end_time, is_delete, create_time) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", data.Title, data.Type, data.DisplayImg, data.BeginTime, data.EndTime, 0, now).Scan(&newArticleID)

	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false
	}

	_, err = tx.Exec("INSERT INTO article_cont(article_id, cont) VALUES($1, $2)", newArticleID, data.Cont)
	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		utils.ErrLog(3, err)
		return false
	}
	return true
}
