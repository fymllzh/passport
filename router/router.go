package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/app/sso"
	"github.com/wuzehv/passport/app/svc"
)

func construct(router *gin.Engine) {
	router.GET("/", sso.Index)
	router.GET("/sso/index", sso.Index)

	router.POST("/sso/login", sso.Login)
	router.POST("/sso/logout", sso.Logout)

	router.POST("/svc/userinfo", svc.Userinfo)
	router.POST("/svc/session", svc.Session)
}
