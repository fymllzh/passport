package sso

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var token = "token_str"

type User struct {
	Name string `json:"name"`
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
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}

// Session 获取session
func Session(c *gin.Context) {
	t:= c.Query("token")
	if t != token {
		fmt.Println("token not match")
		c.JSON(http.StatusInternalServerError, User{})
		return
	}
	s := User{Name: "wzh"}
	c.JSON(http.StatusOK, s)
}

// Auth 检测授权信息
// 客户端每次请求之前都需要先检测授权
func Auth(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
}
