package record

import (
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
	"time"
)

type Record struct {
	base.Model
	UserId    uint   `gorm:"index;not null"`
	ClientId  uint   `gorm:"index;not null"`
	Type      uint   `gorm:"not null;type:tinyint unsigned"`
	IpAddr    string `gorm:"not null;;default:'';type:varchar(255)"`
	UserAgent string `gorm:"not null;default:''"`
}

// 最大失败次数
const LoginFailMaxNum = 5

// 查询区间
const LoginFailDuration = 5 * time.Minute

const (
	// 登录失败
	TypeFail = iota + 1
	// 登录成功
	TypeSuccess
	// 其他(例如：登录错误次数过多导致的重试)
	TypeOther
)

func (u *Record) Base() {}

// FailNumOut 根据指定时间内的错误次数来判断是否可以继续登录
func FailNumOut() bool {
	t := time.Now().Add(-LoginFailDuration)

	var c int64
	db.Db.Model(&Record{}).Where("created_at >= ? and type = ?", t, TypeFail).Count(&c)
	return c >= LoginFailMaxNum
}
