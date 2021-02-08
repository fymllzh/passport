package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/app/sso"
)

func construct(router *gin.Engine) {
	router.GET("/", sso.Index)
	router.GET("/sso/index", sso.Index)
	router.POST("/sso/login", sso.Login)
	router.GET("/sso/logout", sso.Logout)
	router.GET("/sso/auth", sso.Auth)
}
