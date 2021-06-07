package client

import (
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
)

type Client struct {
	base.Model
	Domain   string `gorm:"unique;not null"`
	Callback string `gorm:"index;not null"`
	Secret   string `gorm:"not null"`
	Status   uint   `gorm:"not null;"`
}

func (c *Client) Base() {}

func (c *Client) GetByDomain(domain string) {
	db.Db.Where("domain = ?", domain).First(c)
}
