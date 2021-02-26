package sso

import (
	"fmt"
	"github.com/gin-gonic/gin"
	token2 "github.com/wuzehv/passport/model/token"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"net/http"
	"time"
)

const (
	CallbackKey = "callback"
	CookieKey   = "flag"
	TokenKey    = "token"
)

type User struct {
	Name string `json:"name"`
}

type Response struct {
	Success bool
	Message string
	Data    User
}

func Index(c *gin.Context) {
	callback := c.Query(CallbackKey)
	token, err := c.Cookie(CookieKey)
	if err == nil {
		callback += "?" + TokenKey + "=" + token

		c.HTML(http.StatusOK, "redirect.html", gin.H{
			"callback": callback,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"callback": callback,
		})
	}
}

func Login(c *gin.Context) {
	callback := c.PostForm(CallbackKey)
	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	var u user.User
	db.Db.Where("email = ?", name).First(&u)

	if !util.VerifyPassword(u.Password, passwd) {
		c.JSON(http.StatusOK, Response{Success: false, Message: "用户名或密码错误"})
		return
	}

	// 随便生成一个token
	token := fmt.Sprintf("%d", time.Now().UnixNano())

	// 设置会话
	c.SetCookie(CookieKey, token, 86400, "/", util.ENV("", "domain"), false, true)

	// 持久化
	db.Db.Create(&token2.Token{UserId: u.Id, Token: token})

	// callback
	callback += "?" + TokenKey + "=" + token

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(TokenKey, "false", -1, "/", util.ENV("", "domain"), false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}

// Session 获取session
func Session(c *gin.Context) {
	t := c.Query(TokenKey)

	var token token2.Token
	db.Db.Where("token = ?", t).First(&token)
	if token.Id == 0 {
		c.JSON(http.StatusOK, Response{Success: false, Message: "token not exists"})
		return
	}

	// todo 检测是否过期
	if false {
		c.JSON(http.StatusOK, Response{Success: false, Message: "token expired"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "success", Data: User{Name: "wzh"}})
}

// Auth 检测授权信息
// 客户端每次请求之前都需要先检测授权
func Auth(c *gin.Context) {
	t := c.Query(TokenKey)

	var token token2.Token
	db.Db.Where("token = ?", t).First(&token)
	if token.Id == 0 {
		c.JSON(http.StatusOK, Response{Success: false, Message: "token not exists"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "success"})
}
