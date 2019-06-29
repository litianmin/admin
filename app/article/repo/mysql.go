package repo

import (
	"admin/app/article/entity"
	"admin/utils"
	"database/sql"
	"log"
)

// MysqlRepo 定义结构体
type MysqlRepo struct {
	Conn *sql.DB
}

// NewMysqlRepo 初始化 Repo
func NewMysqlRepo(conn *sql.DB) *MysqlRepo {
	return &MysqlRepo{conn}
}

// NewArticle 新增文章
func (r *MysqlRepo) NewArticle(data *entity.NewArticle) (IsSuccess bool, Article *entity.ArticleBaseInfo) {
	now := utils.NowFormatUnix()
	res, err := r.Conn.Exec("INSERT INTO official_article(title, article_type, display_img, begin_time, end_time, is_delete, create_time) VALUES(?, ?, ?, ?, ?, ?, ?)", data.Title, data.Type, data.DisplayImg, data.BeginTime, data.EndTime, 0, now)

	if err != nil {
		log.Println(err)
		return false, nil
	}

	newArticleID, _ := res.LastInsertId()

	res, err = r.Conn.Exec("INSERT INTO official_article_cont(article_id, cont) VALUES(?, ?)", newArticleID, data.Cont)
	if err != nil {
		log.Println(err)
		return false, nil
	}

	article := entity.ArticleBaseInfo{
		ArticleID:  uint64(newArticleID),
		Title:      data.Title,
		Type:       data.Type,
		BeginTime:  data.BeginTime,
		EndTime:    data.EndTime,
		DisplayImg: data.DisplayImg,
	}
	return true, &article
}
