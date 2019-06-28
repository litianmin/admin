package entity

// AddrObj 地址对象
type AddrObj struct {
	Name string  `json:"name" bson:"name"` // 地址名称
	Lng  float64 `json:"lng" bson:"lng"`   // 经度
	Lat  float64 `json:"lat" bson:"lat"`   // 纬度
	Addr string  `json:"addr" bson:"addr"` // 详细地址
}

// NewActivity 新的活动
type NewActivity struct {
	Title       string  `json:"title"`
	Type        uint8   `json:"type"`
	BeginTime   int64   `json:"beginTIme"`
	EndTime     int64   `json:"endTime"`
	Venue       AddrObj `json:"venue"`
	DisplayImg  string  `json:"displayImg"`
	RecruitNumb uint8   `json:"recruitNumb"`
	Cont        string  `json:"cont"`
}

// ActivityBaseInfo 队伍基本信息
type ActivityBaseInfo struct {
	ActivityID    uint64  `json:"activityID" bson:"activityID"`
	Title         string  `json:"title" bson:"title"`
	Type          uint8   `json:"type" bson:"type"`
	BeginTime     int64   `json:"beginTIme" bson:"beginTime"`
	EndTime       int64   `json:"endTime" bson:"endTime"`
	Venue         AddrObj `json:"venue" bson:"venue"`
	DisplayImg    string  `json:"displayImg" bson:"displayImg"`
	RecruitNumb   uint8   `json:"recruitNumb" bson:"recruitNumb"`
	Distance      float64 `json:"distance" bson:"distance"`
	RecruitStatus uint8   `json:"recruitStatus" bson:"recruitStatus"`
}
