package client

import (
	"errors"
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/service/db"
	"gorm.io/gorm"
)

type Client struct {
	base.Model
	Domain   string `gorm:"unique;not null"`
	Callback string `gorm:"index;not null"`
	Secret   string `gorm:"not null;type:varchar(255)"`
	Status   uint   `gorm:"not null;type:tinyint unsigned" json:"-"`
}

func (c *Client) Base() {}

func (c *Client) GetByDomain(domain string) error {
	if err := db.Db.Where("domain = ?", domain).First(c).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// sql错误
		return err
	}

	return nil
}
