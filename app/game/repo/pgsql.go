package repo

import (
	"admin/app/game/entity"
	"admin/utils"
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// PgRepo 定义结构体
type PgRepo struct {
	Conn *sql.DB
}

// NewPgRepo 初始化 Repo
func NewPgRepo(conn *sql.DB) *PgRepo {
	return &PgRepo{conn}
}

// CreateGame 增加游戏
func (pg *PgRepo) CreateGame(g *entity.NewGame) bool {
	var gameTp int
	switch g.PlatForm {
	case "pc":
		gameTp = 0
	case "mobile":
		gameTp = 0
	}

	now := utils.NowFormatUnix()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := pg.Conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		utils.ErrLog(3, err)
		return false
	}

	var gameID uint64
	// 首先往game_base_info表插入游戏基本信息
	err = tx.QueryRow("INSERT INTO game_base_info (g_name, g_type, g_logo, g_logo_mini, brief_desc, create_time, update_time) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", g.Name, gameTp, g.LogoOrigin, g.LogoMini, g.BriefDesc, now, now).Scan(&gameID)
	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false
	}

	// 插入party的图片
	stmt, err := tx.Prepare(pq.CopyIn("game_display_cont", "g_id", "c_link", "c_type", "position_sort", "create_time"))
	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false
	}

	for index, item := range g.DisplayImgList {
		_, err = stmt.Exec(gameID, item, 1, index, now)
		if err != nil {
			tx.Rollback()
			utils.ErrLog(3, err)
			return false
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false
	}

	_, err = tx.Exec("INSERT INTO game_download_info(g_id, download_link, create_time, update_time) VALUES($1, $2, $3, $4)", gameID, g.DownloadLink, now, now)
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
