package util

import (
	"github.com/go-ini/ini"
)

func ENV(section, key string) string {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	sectionInfo := config.Section(section)
	val, err := sectionInfo.GetKey(key)

	if err != nil {
		panic(err)
	}

	return val.String()
}
