// 平台内部接口

package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
)

func Index(c *gin.Context) {
	domain := c.Query(util.Domain)
	jump := c.Query(util.Jump)

	// 根据token解析出用户信息
	token, err := c.Cookie(util.CookieKey)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"domain": domain,
			"jump": jump,
		})
		return
	}

	uid, err := strconv.Atoi(token[32:])
	if err != nil {
		log.Printf("parse token error: %s\n", err)
		c.AbortWithStatusJSON(http.StatusOK, util.SystemError.Msg(nil))
	}

	var cl client.Client
	cl.GetByDomain(c.Query(util.Domain))
	if cl.Id == 0 {
		c.AbortWithStatusJSON(http.StatusOK, util.ClientNotExists.Msg(nil))
	}

	s := session.NewSession(uint(uid), cl.Id)

	callback := cl.Callback
	callback += "?" + util.TokenKey + "=" + s.Token
	callback += "&" + util.Jump + "=" + url2.QueryEscape(jump)

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Login(c *gin.Context) {
	j, ok := c.Get(util.Jump)
	if !ok {
		// 没有参数
	}

	jump := j.(string)

	j, ok = c.Get(util.Client)
	if !ok {
		// 没有参数
	}

	cl := j.(*client.Client)

	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	var u user.User
	db.Db.Where("email = ?", name).First(&u)

	if !util.VerifyPassword(u.Password, passwd) {
		c.AbortWithStatusJSON(http.StatusOK, util.UsernamePasswdNotMatch.Msg(nil))
	}

	// 初始化token
	md5token := util.GenToken()

	// 设置会话为浏览器关闭即失效
	token := md5token + strconv.FormatUint(uint64(u.Id), 10)
	c.SetCookie(util.CookieKey, token, 0, "/", "", false, true)

	// 重置所有客户端session状态
	session.LogoutAll(u.Id)

	// 持久化
	s := session.NewSession(u.Id, cl.Id)

	// callback
	callback := cl.Callback
	callback += "?" + util.TokenKey + "=" + s.Token
	callback += "&" + util.Jump + "=" + url2.QueryEscape(jump)

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"callback": callback,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(util.TokenKey, "false", -1, "/", util.ENV("", "domain"), false, true)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}
