package base

import "time"

type Model struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Base interface {
	Base()
}
