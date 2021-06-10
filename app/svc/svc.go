// 外部客户端接口

package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
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

	if code := commonCheck(c, &s); code != util.Success {
		c.AbortWithStatusJSON(http.StatusOK, code.Msg(nil))
		return
	}

	// 登录状态
	if s.Status != session.StatusLogin {
		c.AbortWithStatusJSON(http.StatusOK, util.SessionStatusNotLogin.Msg(nil))
		return
	}

	var u user.User
	db.Db.First(&u, s.UserId)
	c.JSON(http.StatusOK, util.Success.Msg(u))
}

// Session 客户端回调确认接口
// 更新session状态为已登录
func Session(c *gin.Context) {
	t := c.Query(util.TokenKey)

	var s session.Session
	db.Db.Where("token = ?", t).First(&s)

	if code := commonCheck(c, &s); code != util.Success {
		c.AbortWithStatusJSON(http.StatusOK, code.Msg(nil))
		return
	}

	if s.Status != session.StatusInit {
		c.AbortWithStatusJSON(http.StatusOK, util.SystemError.Msg(nil))
		return
	}

	// 更新session状态
	s.Status = session.StatusLogin
	db.Db.Save(&s)

	c.JSON(http.StatusOK, util.Success.Msg(nil))
}

func commonCheck(c *gin.Context, s *session.Session) util.Code {
	if s.Id == 0 {
		return util.TokenNotExists
	}

	var u user.User
	db.Db.First(&u, s.UserId)

	if u.Id == 0 || u.Status != base.StatusNormal {
		return util.UserDisabled
	}

	j, _ := c.Get(util.Client)
	cl := j.(*client.Client)

	if cl.Id == 0 || cl.Status != base.StatusNormal {
		return util.ClientDisabled
	}

	// 客户端和session不匹配
	if cl.Id != s.ClientId {
		return util.SystemError
	}

	// 过期检测
	if time.Now().After(s.ExpireTime) {
		return util.SessionExpired
	}

	return util.Success
}
