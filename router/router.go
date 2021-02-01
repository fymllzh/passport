package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/app/sso"
)

func construct(router *gin.Engine) {
	router.GET("/", sso.Login)
	router.GET("/sso/login", sso.Login)
	router.GET("/sso/logout", sso.Logout)
}
