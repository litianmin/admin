package entity

// NewGame 新的游戏结构体
type NewGame struct {
	Name           string   `json:"name" valid:"required"`
	PlatForm       string   `json:"platform" valid:"required"`
	LogoOrigin     string   `json:"logoOrigin" valid:"required"`
	LogoMini       string   `json:"logoMini" valid:"required"`
	DisplayImgList []string `json:"displayImgList" valid:"required"`
	DownloadLink   string   `json:"downloadLink" valid:"required"`
	BriefDesc      string   `json:"briefDesc" valid:"required"`
}
