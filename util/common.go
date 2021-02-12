package util

import (
	"github.com/go-ini/ini"
)

func ENV(section, key string) string {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	val, err := config.Section(section).GetKey(key)

	if err != nil {
		panic(err)
	}

	return val.String()
}
