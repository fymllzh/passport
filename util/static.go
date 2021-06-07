package util

type Code int

// 状态码
const (
	Success Code = iota
	// 外部状态码
	UserDisabled
	TokenNotExists
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
	UserDisabled:           "user disabled",
	TokenNotExists:         "token not exists",
	SessionExpired:         "session expired",
	SessionStatusNotLogin:  "session status not login",
	ClientNotExists:        "client not exists",
	ClientDisabled:         "client disabled",
	SystemError:            "system error",
	UsernamePasswdNotMatch: "用户名密码错误",
}

// 通用key
const (
	Domain    = "domain"
	Jump      = "jump"
	CookieKey = "flag"
	TokenKey  = "token"
	Client    = "client"
)

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
