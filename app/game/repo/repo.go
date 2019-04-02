package repo

import (
	"admin/app/game/entity"
	"admin/utils"
	"database/sql"
	"fmt"
)

// Repo 定义结构体
type Repo struct {
	Conn *sql.DB
}

// NewRepo 初始化 Repo
func NewRepo(conn *sql.DB) *Repo {
	return &Repo{conn}
}

// CreateGame 增加游戏
func (r *Repo) CreateGame(g *entity.NewGame) bool {
	var gameTp string
	switch g.PlatForm {
	case "pc":
		gameTp = "0"
	case "mobile":
		gameTp = "1"
	}

	now := utils.NowFormatToDate()

	// 首先往game_base_info表插入游戏基本信息
	res, err := r.Conn.Exec("INSERT INTO game_base_info(g_name, g_type, g_logo, g_logo_mini, brief_desc, create_time, update_time) VALUES(?, ?, ?, ?, ?, ?, ?)", g.Name, gameTp, g.LogoOrigin, g.LogoMini, g.BriefDesc, now, now)

	if err != nil {
		fmt.Println(err)
		return false
	}

	gameID, _ := res.LastInsertId()

	// 得到了gameID 之后，往 game_display_cont
	stmt, err := r.Conn.Prepare("INSERT INTO game_display_cont(g_id, c_link, c_type, position_sort, create_time) VALUES(?, ?, ?, ?, ?)")
	for i := 0; i < len(g.DisplayImgList); i++ {
		stmt.Exec(gameID, g.DisplayImgList[i], "1", i, now)
	}

	defer stmt.Close()

	return true
}
