// 外部客户端接口

package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"net/http"
	"time"
)

// UserInfo 获取用户信息
// 客户端业务代码执行之前，需要调用该接口获取用户信息
func Userinfo(c *gin.Context) {
	t := c.Query(util.TokenKey)

	var s session.Session
	db.Db.Where("token = ?", t).First(&s)

	if code := commonCheck(&s); code != util.Success {
		c.JSON(http.StatusOK, code.Msg(nil))
		return
	}

	// 登录状态
	if s.Status != session.StatusLogin {
		c.JSON(http.StatusOK, util.SessionStatusNotLogin.Msg(nil))
		return
	}

	c.JSON(http.StatusOK, util.Success.Msg(nil))
}

// Session 客户端回调确认接口
// 更新session状态为已登录
func Session(c *gin.Context) {
	t := c.Query(util.TokenKey)

	var s session.Session
	db.Db.Where("token = ?", t).First(&s)

	if code := commonCheck(&s); code != util.Success {
		c.JSON(http.StatusOK, code.Msg(nil))
		return
	}

	if s.Status != session.StatusInit {
		c.JSON(http.StatusOK, util.SystemError.Msg(nil))
		return
	}

	// 更新session状态
	s.Status = session.StatusLogin
	db.Db.Save(&s)

	c.JSON(http.StatusOK, util.Success.Msg(nil))
}

func commonCheck(s *session.Session) util.Code {
	if s.Id == 0 {
		return util.TokenNotExists
	}

	// 过期检测
	if time.Now().After(s.ExpireTime) {
		return util.SessionExpired
	}

	// 客户端检测

	return util.Success
}
