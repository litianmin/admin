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

// NewActivity 创建新的活动
func (pg *PgRepo) NewActivity(data *entity.NewActivity) (IsSuccess bool, NewActivityID int64) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := pg.Conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		utils.ErrLog(3, err)
		return false, 0
	}

	now := utils.NowFormatUnix()

	// 插入活动基本数据(因为geo数据类型问题，需要这样赋值)
	sql := fmt.Sprintf("INSERT INTO activity (sponsor, title, act_type, display_img, site_lng, site_lat, site_geog, site_name, site_detail, begin_time, end_time, recruit_numb, is_official, apply_status, create_time) VALUES ($1, $2, $3, $4, $5, $6, 'SRID=4326;POINT(%f %f)', $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id", data.Venue.Lng, data.Venue.Lat)

	var newID int64

	err = tx.QueryRow(sql, 0, data.Title, data.Type, data.DisplayImg, data.Venue.Lng, data.Venue.Lat, data.Venue.Name, data.Venue.Addr, data.BeginTime, data.EndTime, data.RecruitNumb, 1, 2, now).Scan(&newID)

	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false, 0
	}

	// 查询活动详细内容
	_, err = tx.Exec("INSERT INTO activity_cont (activity_id, cont) VALUES ($1, $2)", newID, data.Cont)
	if err != nil {
		tx.Rollback()
		utils.ErrLog(3, err)
		return false, 0
	}

	tx.Commit()

	return true, newID
}
