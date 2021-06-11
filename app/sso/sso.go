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
	tmp, _ := c.Get(util.Client)
	cl := tmp.(*client.Client)

	tmp, _ = c.Get(util.Jump)
	jump := tmp.(string)

	u, _ := c.Get(util.Uid)
	uid := u.(uint)

	if uid == 0 {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"domain": cl.Domain,
			"jump":   jump,
		})
		return
	}

	commonDeal(c, uid, jump)
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

	commonDeal(c, u.Id, jump)
}

func commonDeal(c *gin.Context, userId uint, jump string) {
	tmp, _ := c.Get(util.Client)
	cl := tmp.(*client.Client)

	// 持久化
	s := session.NewSession(userId, cl.Id)

	callback := cl.Callback
	callback += "?" + util.TokenKey + "=" + s.Token
	callback += "&" + util.Jump + "=" + url.QueryEscape(jump)

	tmp, _ = c.Get(util.Sso)
	isSso := tmp.(bool)

	if isSso {
		c.HTML(http.StatusOK, "redirect.html", gin.H{
			"callback": callback,
		})
	} else {
		// 如果不是sso，跳转到首页
		c.Redirect(http.StatusTemporaryRedirect, "/admin/index/index")
	}
}

func Logout(c *gin.Context) {
	u, _ := c.Get(util.Uid)
	uid := u.(uint)
	session.LogoutAll(uid)

	c.SetCookie(util.CookieKey, "false", -1, "/", "", false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}
