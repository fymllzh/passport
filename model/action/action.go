package action

import (
	"github.com/wuzehv/passport/model/base"
)

type Action struct {
	base.Model
	Url    string `gorm:"unique" json:"url"`
	Remark string `gorm:"not null" json:"remark"`
}

func (u *Action) Base() {}
