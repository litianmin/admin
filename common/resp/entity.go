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
		Msg:  "Account Or Pwd Err!",
	}

	// ParamsErr 参数错误
	ParamsErr = GeneralResp{
		Code: 40001,
		Msg:  "Params Err!",
	}

	// TokenWithout 未登录
	TokenWithout = GeneralResp{
		Code: 40002,
		Msg:  "Without Token!",
	}

	// TokenIllegal 非法token
	TokenIllegal = GeneralResp{
		Code: 40003,
		Msg:  "Token Is Illegal!",
	}

	// TokenExpired token过期
	TokenExpired = GeneralResp{
		Code: 40004,
		Msg:  "Token Is Expired!",
	}
)
