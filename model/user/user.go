package user

import "github.com/wuzehv/passport/model/base"

type User struct {
	base.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null;" json:"-"`
	Status   uint   `gorm:"not null;" json:"-"`
}

func (u *User) Base() {}
