package entity

// LoginAuth 登陆验证
type LoginAuth struct {
	UserName string `json:"u_name"`
	Pwd      string `json:"pwd"`
}
