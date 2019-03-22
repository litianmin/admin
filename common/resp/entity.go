package resp

// GeneralResp 定义一般回复实体
type GeneralResp struct {
	Code uint64
	Msg  interface{}
}

var (
	// AccountPwdErr 账号密码错误
	AccountPwdErr = GeneralResp{
		Code: 40000,
		Msg:  "账号密码错误",
	}
)
