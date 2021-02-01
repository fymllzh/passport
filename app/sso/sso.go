package sso

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Logout(c *gin.Context) {
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}
