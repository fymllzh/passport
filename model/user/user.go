package user

import "github.com/wuzehv/passport/model/base"

type User struct {
	base.Model
	Email    string `gorm:"unique"`
	Password string
	Status   uint
}

const (
	// 正常
	StatusNormal = iota
	// 已禁用
	StatusDisabled
)

func (u *User) Base() {}
