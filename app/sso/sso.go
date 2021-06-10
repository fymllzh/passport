// 平台内部接口

package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/util"
	"net/http"
	"net/url"
	"strconv"
)

func Index(c *gin.Context) {
	domain := c.Query(util.Domain)
	jump := c.Query(util.Jump)

	u, _ := c.Get(util.Uid)
	uid := u.(int)

	if uid == 0 {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"domain": domain,
			"jump":   jump,
		})
		return
	}

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": commonDeal(c, uint(uid), jump),
	})
}

func Login(c *gin.Context) {
	j, _ := c.Get(util.Jump)
	jump := j.(string)

	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	var u user.User
	u.GetByEmail(name)

	if !util.VerifyPassword(u.Password, passwd) {
		c.AbortWithStatusJSON(http.StatusOK, util.UsernamePasswdNotMatch.Msg(nil))
		return
	}

	// 设置会话为浏览器关闭即失效
	token := util.GenToken() + strconv.FormatUint(uint64(u.Id), 10)
	c.SetCookie(util.CookieKey, token, 0, "/", "", false, true)

	// 重置所有客户端session状态
	session.LogoutAll(u.Id)

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": commonDeal(c, uint(u.Id), jump),
	})
}

func commonDeal(c *gin.Context, userId uint, jump string) (callback string) {
	j, _ := c.Get(util.Client)
	cl := j.(*client.Client)

	// 持久化
	s := session.NewSession(userId, cl.Id)

	callback = cl.Callback
	callback += "?" + util.TokenKey + "=" + s.Token
	callback += "&" + util.Jump + "=" + url.QueryEscape(jump)

	return
}

func Logout(c *gin.Context) {
	c.SetCookie(util.TokenKey, "false", -1, "/", util.ENV("", "domain"), false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}
