package resp

// GeneralResp 定义一般回复实体
type GeneralResp struct {
	Code uint64      `json:"code"`
	Msg  interface{} `json:"msg"`
}

var (
	// AccountPwdErr 账号密码错误
	AccountPwdErr = GeneralResp{
		Code: 40000,
		Msg:  "账号密码错误",
	}

	// ParamsErr 参数错误
	ParamsErr = GeneralResp{
		Code: 40001,
		Msg:  "参数错误",
	}
)
