// 外部客户端接口

package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"net/http"
)

// Session 客户端回调确认接口
// 更新session状态为已登录
func Session(c *gin.Context) {
	tmp, _ := c.Get(util.Session)
	s := tmp.(session.Session)

	if s.Status != session.StatusInit {
		c.AbortWithStatusJSON(http.StatusOK, util.SystemError.Msg(nil))
		return
	}

	// 更新session状态
	s.Status = session.StatusLogin
	db.Db.Save(&s)

	c.JSON(http.StatusOK, util.Success.Msg(nil))
}

// UserInfo 获取用户信息
// 客户端业务代码执行之前，需要调用该接口获取用户信息
func Userinfo(c *gin.Context) {
	tmp, _ := c.Get(util.Session)
	s := tmp.(session.Session)

	// 登录状态
	if s.Status != session.StatusLogin {
		c.AbortWithStatusJSON(http.StatusOK, util.SessionStatusNotLogin.Msg(nil))
		return
	}

	tmp, _ = c.Get(util.User)
	u := tmp.(user.User)
	c.JSON(http.StatusOK, util.Success.Msg(u))
}
