// 平台内部接口

package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	callback := c.Query(util.CallbackKey)

	// 根据token解析出用户信息
	token, err := c.Cookie(util.CookieKey)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"callback": callback,
		})

		return
	}

	uid, _ := strconv.Atoi(token[32:])
	s := session.NewSession(uint(uid), 0)

	callback += "?" + util.TokenKey + "=" + s.Token

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Login(c *gin.Context) {
	callback := c.PostForm(util.CallbackKey)
	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	var u user.User
	db.Db.Where("email = ?", name).First(&u)

	if !util.VerifyPassword(u.Password, passwd) {
		c.JSON(http.StatusOK, util.UsernamePasswdNotMatch.Msg(nil))
		return
	}

	// 初始化token
	md5token := util.GenToken()

	// 设置会话为浏览器关闭即失效
	token := md5token + strconv.FormatUint(uint64(u.Id), 10)
	c.SetCookie(util.CookieKey, token, 0, "/", "", false, true)

	// 重置所有客户端session状态
	session.LogoutAll(u.Id)

	// 持久化
	s := session.NewSession(u.Id, 0)

	// callback
	callback += "?" + util.TokenKey + "=" + s.Token

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(util.TokenKey, "false", -1, "/", util.ENV("", "domain"), false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}
