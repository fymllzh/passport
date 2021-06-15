package util

type Code int

// 状态码
const (
	Success Code = iota
	// 外部状态码
	ParamsError
	SignatureError
	UserDisabled
	TokenNotExists
	TokenParseError
	SessionExpired
	SessionStatusNotLogin
	ClientNotExists
	ClientDisabled
	SystemError
	// 内部状态码
	UsernamePasswdNotMatch
)

var errors = [...]string{
	Success:                "success",
	ParamsError:            "params error",
	SignatureError:         "signature error",
	UserDisabled:           "user disabled",
	TokenNotExists:         "token not exists",
	TokenParseError:        "token parse exists",
	SessionExpired:         "session expired",
	SessionStatusNotLogin:  "session status not login",
	ClientNotExists:        "client not exists",
	ClientDisabled:         "client disabled",
	SystemError:            "system error",
	UsernamePasswdNotMatch: "用户名密码错误",
}

// 通用key
const (
	Domain     = "domain"
	Jump       = "jump"
	CookieFlag = "flag"
	Token      = "token"
	Client     = "client"
	Uid        = "uid"
	Sso        = "sso"
	Session    = "session"
	User       = "user"
	Timestamp  = "timestamp"
	Sign       = "sign"
	Secret     = "secret"
)

type SvcRequest struct {
	Token     string `form:"token"`
	Domain    string `form:"domain"`
	Timestamp string `form:"timestamp"`
	Sign      string `form:"sign"`
}

// 响应结构体
type Response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c Code) Msg(data interface{}) Response {
	return Response{
		Code:    c,
		Message: errors[c],
		Data:    data,
	}
}

func (c Code) Error() string {
	return errors[c]
}
