package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/app/admin/index"
	"github.com/wuzehv/passport/app/sso"
	"github.com/wuzehv/passport/app/svc"
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func construct(router *gin.Engine) {
	// sso主页
	ssoIn := router.Group("/")
	ssoIn.Use(ssoBase())
	{
		ssoIn.GET("/", sso.Index)
		ssoIn.GET("/sso/index", sso.Index)

		ssoIn.POST("/sso/login", sso.Login)
		ssoIn.GET("/sso/logout", sso.Logout)
	}

	// admin内部
	admin := router.Group("/admin")
	admin.Use(adminBase())
	{
		admin.GET("/index/index", index.Index)
		admin.GET("/index/test", index.Test)
	}

	// 对外接口
	scvIn := router.Group("/svc")
	scvIn.Use(svcBase())
	{
		scvIn.POST("/userinfo", svc.Userinfo)
		scvIn.POST("/session", svc.Session)
	}
}

// sso中心页面入口
func ssoBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := c.Query(util.Domain)
		var cl client.Client
		cl.GetByDomain(domain)

		c.Set(util.Sso, true)
		if cl.Id > 0 && cl.Status != base.StatusNormal {
			c.AbortWithError(http.StatusForbidden, util.ClientDisabled)
			return
		}

		if cl.Id == 0 {
			c.Set(util.Sso, false)
		}

		c.Set(util.Client, &cl)
		c.Set(util.Jump, c.Query(util.Jump))
		c.Set(util.Uid, uint(0))

		// 根据token解析出用户信息
		token, err := c.Cookie(util.CookieFlag)
		if err != nil {
			return
		}

		uid, err := strconv.Atoi(token[32:])
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, util.TokenParseError)
			return
		}

		c.Set(util.Uid, uint(uid))
	}
}

// admin页面
func adminBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(util.CookieFlag)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		uid, err := strconv.Atoi(token[32:])
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		var u user.User
		db.Db.First(&u, uid)
		if u.Id == 0 || u.Status != base.StatusNormal {
			fmt.Fprintln(c.Writer, util.UserDisabled.Msg(nil))
			return
		}

		// 判断登录是否过期
		if u.Token != token || time.Now().After(u.ExpireTime) {
			// 显式的删除cookie
			c.SetCookie(util.CookieFlag, "false", -1, "/", "", false, true)

			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		c.Set(util.User, &u)
	}
}

// svcBase svc调用入口，校验token
func svcBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		var res util.SvcRequest
		if c.ShouldBind(&res) != nil {
			c.AbortWithStatusJSON(http.StatusOK, util.ParamsError.Msg(nil))
			return
		}

		domain := res.Domain
		domain, _ = url.QueryUnescape(domain)

		var cl client.Client
		cl.GetByDomain(domain)

		if cl.Id == 0 || cl.Status != base.StatusNormal {
			c.AbortWithStatusJSON(http.StatusOK, util.ClientDisabled.Msg(nil))
			return
		}

		m := make(map[string]string)
		m[util.Token] = res.Token
		m[util.Timestamp] = res.Timestamp
		m[util.Domain] = res.Domain
		if util.GenSign(m, cl.Secret) != res.Sign {
			c.AbortWithStatusJSON(http.StatusOK, util.SignatureError.Msg(nil))
			return
		}

		t := res.Token

		var s session.Session
		s.GetByToken(t)

		if s.Id == 0 {
			c.AbortWithStatusJSON(http.StatusOK, util.TokenNotExists.Msg(nil))
			return
		}

		var u user.User
		db.Db.First(&u, s.UserId)

		if u.Id == 0 || u.Status != base.StatusNormal {
			c.AbortWithStatusJSON(http.StatusOK, util.UserDisabled.Msg(nil))
			return
		}

		// 客户端和session不匹配
		if cl.Id != s.ClientId {
			c.AbortWithStatusJSON(http.StatusOK, util.SystemError.Msg(nil))
			return
		}

		// 过期检测
		if time.Now().After(s.ExpireTime) {
			c.AbortWithStatusJSON(http.StatusOK, util.SessionExpired.Msg(nil))
			return
		}

		c.Set(util.Session, &s)
		c.Set(util.User, &u)
	}
}
