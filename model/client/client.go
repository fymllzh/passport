package client

import (
	"github.com/wuzehv/passport/model/base"
)

type Client struct {
	base.Model
	Domain   string `gorm:"unique;not null"`
	Callback string `gorm:"index;not null"`
	Secret   string `gorm:"not null"`
	Status   uint   `gorm:"not null;"`
}

const (
	// 正常
	StatusNormal = iota
	// 已禁用
	StatusDisabled
)

func (c *Client) Base() {}
