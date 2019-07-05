package repo

import (
	"admin/app/activity/entity"
	"admin/utils"
	"context"
	"database/sql"
	"fmt"
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

// type NewActivity struct {
// 	Title       string  `json:"title"`
// 	Type        uint8   `json:"type"`
// 	BeginTime   int64   `json:"beginTIme"`
// 	EndTime     int64   `json:"endTime"`
// 	Venue       AddrObj `json:"venue"`
// 	DisplayImg  string  `json:"displayImg"`
// 	RecruitNumb uint8   `json:"recruitNumb"`
// 	Cont        string  `json:"cont"`
// }

// NewActivity 创建新的活动
func (pg *PgRepo) NewActivity(data *entity.NewActivity) (IsSuccess bool, NewActivityID int64) {

	// now := utils.NowFormatUnix()

	// // pg.Conn.BeginTx

	// res, err := pg.Conn.Exec("INSERT INTO activity(sponsor, title, act_type, display_img, site_lng, site_lat, site_geog, site_name, site_detail, begin_time, end_time, recruit_numb, is_official, apply_status, create_time) VALUES(?, ?, ?, ?, ?, ?, 'SRID=4326;POINT(? ?)', ?, ?, ?, ?, ?, ?, ?, ?)", 0, data.Title, data.Type, data.DisplayImg, data.Venue.Lng, data.Venue.Lat, data.Venue.Lng, data.Venue.Lat, data.Venue.Name, data.Venue.Addr, data.BeginTime, data.EndTime, data.RecruitNumb, 1, 2, now)

	// if err != nil {
	// 	utils.ErrLog(3, err)
	// 	return false, 0
	// }
	// newActivityID, _ := res.LastInsertId()

	// fmt.Println(newActivityID)

	// fmt.Println("什么鬼啊")

	// 插入内容

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := pg.Conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		utils.ErrLog(3, err)
		return false, 0
	}

	now := utils.NowFormatUnix()

	fmt.Println(data.Venue.Lng)
	fmt.Println(data.Venue.Lat)

	// 插入数据
	res, err := tx.Exec("INSERT INTO activity (sponsor, title, act_type, display_img, site_lng, site_lat, site_geog, site_name, site_detail, begin_time, end_time, recruit_numb, is_official, apply_status, create_time) VALUES ($1, $2, $3, $4, $5, $6, 'SRID=4326;POINT(113.186702 23.035872)', $7, $8, $9, $10, $11, $12, $13, $14)", 0, data.Title, data.Type, data.DisplayImg, data.Venue.Lng, data.Venue.Lat, data.Venue.Name, data.Venue.Addr, data.BeginTime, data.EndTime, data.RecruitNumb, 1, 2, now)

	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false, 0
	}

	newID, _ := res.LastInsertId()

	tx.Commit()
	fmt.Println(newID)
	return true, 1
}
