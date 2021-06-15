package user

import (
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
	"time"
)

type User struct {
	base.Model
	Email      string    `gorm:"unique" json:"email"`
	Password   string    `gorm:"not null;type:varchar(255)" json:"-"`
	Token      string    `gorm:"unique;not null;default:''" json:"-"`
	ExpireTime time.Time `json:"-"`
}

func (u *User) Base() {}

func (u *User) GetByEmail(email string) {
	db.Db.Where("email = ?", email).First(u)
}
