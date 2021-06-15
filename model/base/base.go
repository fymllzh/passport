package base

import "time"

type Model struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Status    uint      `gorm:"not null;type:tinyint unsigned" json:"-"`
}

const (
	// 正常
	StatusNormal = iota + 1
	// 已禁用
	StatusDisabled
)

type Base interface {
	Base()
}
