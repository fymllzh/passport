// 平台内部接口

package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/login/record"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Index(c *gin.Context) {
	tmp, _ := c.Get(util.Client)
	cl := tmp.(*client.Client)

	tmp, _ = c.Get(util.Jump)
	jump := tmp.(string)

	u, _ := c.Get(util.Uid)
	uid := u.(uint)

	if uid == 0 {
		c.HTML(http.StatusOK, "sso/login", gin.H{
			"domain": cl.Domain,
			"jump":   jump,
		})
		return
	}

	commonDeal(c, cl, uid, jump)
}

func Login(c *gin.Context) {
	j, _ := c.Get(util.Jump)
	jump := j.(string)

	name := c.PostForm("username")
	passwd := c.PostForm("password")

	// 校验密码
	var u user.User
	u.GetByEmail(name)

	tmp, _ := c.Get(util.Client)
	cl := tmp.(*client.Client)

	// 初始化登录信息
	r := record.Record{
		UserId:    u.Id,
		ClientId:  cl.Id,
		IpAddr:    c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	}

	if record.FailNumOut() {
		r.Type = record.TypeOther
		db.Db.Save(&r)
		c.AbortWithStatusJSON(http.StatusOK, util.UsernamePasswdFailNumOut.Msg(nil))
		return
	}

	if !util.VerifyPassword(u.Password, passwd) {
		r.Type = record.TypeFail
		db.Db.Save(&r)
		c.AbortWithStatusJSON(http.StatusOK, util.UsernamePasswdNotMatch.Msg(nil))
		return
	}

	// 初始化token
	token := util.GenToken() + strconv.FormatUint(uint64(u.Id), 10)
	u.Token = token
	exp, _ := time.Parse("2006-01-02 15:04:05", time.Now().Add(session.ExpireTime).Format("2006-01-02")+" 04:00:00")
	u.ExpireTime = exp
	db.Db.Save(&u)
	// 设置会话为浏览器关闭即失效
	c.SetCookie(util.CookieFlag, token, 0, "/", "", false, true)

	// 重置所有客户端session状态
	session.LogoutAll(u.Id)

	r.Type = record.TypeSuccess
	db.Db.Save(&r)

	commonDeal(c, cl, u.Id, jump)
}

func commonDeal(c *gin.Context, cl *client.Client, userId uint, jump string) {
	// 持久化
	s := session.NewSession(userId, cl.Id)

	callbackUrl, err := url.Parse(cl.Callback)
	if err != nil {
		log.Fatal("callback url config error")
	}

	callbackParams := url.Values{}
	callbackParams.Add(util.Token, s.Token)
	callbackParams.Add(util.Jump, jump)

	callbackUrl.RawQuery = callbackParams.Encode()

	tmp, _ := c.Get(util.Sso)
	isSso := tmp.(bool)

	if isSso {
		c.HTML(http.StatusOK, "sso/redirect", gin.H{
			"callback": callbackUrl,
		})
	} else {
		// 如果不是sso，跳转到首页
		c.Redirect(http.StatusMovedPermanently, "/admin/index/index")
	}
}

func Logout(c *gin.Context) {
	u, _ := c.Get(util.Uid)
	uid := u.(uint)
	session.LogoutAll(uid)

	c.SetCookie(util.CookieFlag, "false", -1, "/", "", false, true)
	c.HTML(http.StatusOK, "sso/logout", gin.H{})
}
