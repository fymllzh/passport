package client

import (
	"github.com/wuzehv/passport/model/base"
)

type Client struct {
	base.Model
	Domain   string `gorm:"unique"`
	Callback uint   `gorm:"index"`
	Secret   string
	Status   uint
}

const (
	// 正常
	StatusNormal = iota
	// 已禁用
	StatusDisabled
)

func (c *Client) Base() {}
