package token

import (
	"time"
)

type Token struct {
	Id        uint   `gorm:"primaryKey"`
	Token     string `gorm:"unique"`
	UserId    uint   `gorm:"index"`
	CreatedAt time.Time
}
