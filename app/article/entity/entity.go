package entity

// title: title,
// type: Number(this.Type),
// displayBeginTime,
// displayEndTime,
// displayImg: this.DisplayImgForSave,
// cont: detail,

// NewArticle 新的文章
type NewArticle struct {
	Title      string `json:"title"`
	Type       uint8  `json:"type"`
	BeginTime  int64  `json:"displayBeginTime"`
	EndTime    int64  `json:"displayEndTime"`
	DisplayImg string `json:"displayImg"`
	Cont       string `json:"cont"`
}

// ArticleBaseInfo 文章基本信息
type ArticleBaseInfo struct {
	ArticleID  uint64 `json:"articleID"`
	Title      string `json:"title"`
	Type       uint8  `json:"type"`
	BeginTime  int64  `json:"beginTime"`
	EndTime    int64  `json:"endTime"`
	DisplayImg string `json:"displayImg"`
}
