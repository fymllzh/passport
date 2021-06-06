package util

type Code int

const (
	// 状态码
	Success Code = iota
	// 外部状态码
	TokenNotExists
	SessionExpired
	SessionStatusNotLogin
	ClientNotExists
	SystemError
	// 内部状态码
	UsernamePasswdNotMatch
)

var errors = [...]string{
	Success:               "success",
	TokenNotExists:        "token not exists",
	SessionExpired:        "session expired",
	SessionStatusNotLogin: "session status not login",
	ClientNotExists:       "client not exists",
	SystemError:           "system error",

	UsernamePasswdNotMatch: "用户名密码错误",
}

const (
	// 魔术key
	CallbackKey = "callback"
	CookieKey   = "flag"
	TokenKey    = "token"
)

// 响应结构体
type Response struct {
	Code    Code
	Message string
	Data    interface{}
}

func (c Code) Msg(data interface{}) Response {
	return Response{
		Code:    c,
		Message: errors[c],
		Data:    data,
	}
}
