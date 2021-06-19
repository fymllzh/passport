package base

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/service/db"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Model struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

const (
	// 正常
	StatusNormal = iota + 1
	// 已禁用
	StatusDisabled
)

const (
	DefaultPageSize = 10
	MaxPageSize     = 100
)

type Base interface {
	Base()
}

type Param struct {
	Where string
	Bind  []interface{}
}

type PaginateResponse struct {
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
	Total    int64                    `json:"total"`
	Items    []map[string]interface{} `json:"items"`
}

func Paginate(page, pageSize int, table interface{}, params Param) (*PaginateResponse, error) {
	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > MaxPageSize:
		pageSize = MaxPageSize
	case pageSize <= 0:
		pageSize = DefaultPageSize
	}

	res := PaginateResponse{
		Page:     page,
		PageSize: pageSize,
		Items:    []map[string]interface{}{},
	}

	if params.Where != "" {
		if err := db.Db.Model(table).Where(params.Where, params.Bind...).Count(&res.Total).Error; err != nil {
			return nil, err
		}

		if err := db.Db.Model(table).Scopes(PaginateScopes(page, pageSize)).Where(params.Where, params.Bind...).Find(&res.Items).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Db.Model(table).Count(&res.Total).Error; err != nil {
			return nil, err
		}

		if err := db.Db.Model(table).Scopes(PaginateScopes(page, pageSize)).Find(&res.Items).Error; err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func Paginate2(c *gin.Context, table interface{}, params Param) (*PaginateResponse, error) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	pageSize, _ := strconv.Atoi(c.PostForm("page_size"))
	return Paginate(page, pageSize, table, params)
}

func PaginateScopes(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
