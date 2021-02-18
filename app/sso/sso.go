package sso

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var token = "token_str"

type User struct {
	Name string `json:"name"`
}

type Response struct {
	Success bool
	Message string
	Data    User
}

func Index(c *gin.Context) {
	callback := c.Query("callback")
	token, err := c.Cookie("token")
	if err == nil {
		callback += "?token=" + token

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
	callback := c.PostForm("callback")
	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	if !(name == "w" && passwd == "123456") {
		c.String(http.StatusOK, "信息错误")
		return
	}

	// 设置会话
	c.SetCookie("token", token, 3600, "/", "sso.com", false, true)
	// 持久化token

	// callback
	callback += "?token=" + token

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", token, -1, "/", "sso.com", false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}

// Session 获取session
func Session(c *gin.Context) {
	t := c.Query("token")
	if t != token {
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
	// todo 获取用户信息
	c.JSON(http.StatusOK, Response{Success: false, Message: "test"})
}
