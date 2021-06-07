package base

import "time"

type Model struct {
	Id        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

const (
	// 正常
	StatusNormal = iota
	// 已禁用
	StatusDisabled
)

type Base interface {
	Base()
}
