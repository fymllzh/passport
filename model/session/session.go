package session

import (
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
	"time"
)

type Session struct {
	base.Model
	Token      string    `gorm:"unique;not null"`
	UserId     uint      `gorm:"index;not null"`
	ClientId   uint      `gorm:"index;not null"`
	ExpireTime time.Time `gorm:"not null;"`
	Status     uint      `gorm:"not null;"`
}

const (
	// 初始化状态
	StatusInit = iota + 1
	// 已登录
	StatusLogin
	// 已退出
	StatusLogout
)

const ExpireTime = 24 * time.Hour

func (s *Session) Base() {}

func NewSession(userId, clientId uint) Session {
	s := Session{
		Token:      util.GenToken(),
		UserId:     userId,
		ClientId:   clientId,
		Status:     StatusInit,
		ExpireTime: time.Now().Add(ExpireTime),
	}
	db.Db.Create(&s)

	return s
}

func (s *Session) GetByToken(t string) {
	db.Db.Where("token = ?", t).First(&s)
}

func LogoutAll(userId uint) {
	db.Db.Model(Session{}).Where("user_id = ?", userId).Updates(Session{Status: StatusLogout})
}
