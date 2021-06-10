package user

import (
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
)

type User struct {
	base.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null;" json:"-"`
	Status   uint   `gorm:"not null;" json:"-"`
}

func (u *User) Base() {}

func (u *User) GetByEmail(email string) {
	db.Db.Where("email = ?", email).First(u)
}
