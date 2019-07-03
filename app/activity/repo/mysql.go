package repo

import (
	"admin/app/activity/entity"
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

// NewNewOfficialActivity 创建新的官方活动
func (r *MysqlRepo) NewNewOfficialActivity(data *entity.NewActivity) (IsSuccess bool, Activity *entity.ActivityBaseInfo) {
	now := utils.NowFormatUnix()

	res, err := r.Conn.Exec("INSERT INTO official_activity(sponsor, title, act_type, display_img, meeting_addr_name, meeting_addr_detail, meeting_lng, meeting_lat, begin_time, end_time, recruit_numb, is_official, apply_status, is_delete, create_time) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 1, data.Title, data.Type, data.DisplayImg, data.Venue.Name, data.Venue.Addr, data.Venue.Lng, data.Venue.Lat, data.BeginTime, data.EndTime, data.RecruitNumb, 1, 2, 0, now)

	if err != nil {
		log.Println(err)
		return false, nil
	}

	newActivityID, _ := res.LastInsertId() // 新的活动ID，现在的话，应该怎么做呢，我也不是很清楚耶

	// 往 official_activity_cont 表插入内容
	res, err = r.Conn.Exec("INSERT INTO official_activity_cont(activity_id, cont) VALUES(?, ?)", newActivityID, data.Cont)

	if err != nil {
		log.Println(err)
		return false, nil
	}

	activityBaseInfo := entity.ActivityBaseInfo{
		ActivityID:    uint64(newActivityID),
		Title:         data.Title,
		Type:          data.Type,
		BeginTime:     data.BeginTime,
		EndTime:       data.EndTime,
		Venue:         data.Venue,
		DisplayImg:    data.DisplayImg,
		RecruitNumb:   data.RecruitNumb,
		Distance:      0,
		RecruitStatus: 0,
	}

	return true, &activityBaseInfo
}
