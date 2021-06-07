package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/app/sso"
	"github.com/wuzehv/passport/app/svc"
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/util"
	"net/http"
)

func construct(router *gin.Engine) {
	router.GET("/", sso.Index)
	router.GET("/sso/index", sso.Index)

	router.POST("/sso/login", svcBase(), sso.Login)
	router.POST("/sso/logout", svcBase(), sso.Logout)

	router.POST("/svc/userinfo", svcBase(), svc.Userinfo)
	router.POST("/svc/session", svcBase(), svc.Session)
}

func ssoBase() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// svcBase svc接口入口
func svcBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := c.Query(util.Domain)
		var cl client.Client
		cl.GetByDomain(domain)

		if cl.Id == 0 || cl.Status != base.StatusNormal {
			c.AbortWithStatusJSON(http.StatusOK, util.ClientDisabled.Msg(nil))
		}

		c.Set(util.Client, &cl)
		c.Set(util.Jump, c.Query(util.Jump))
		c.Next()
	}
}
