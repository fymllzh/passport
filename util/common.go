package util

import (
	"github.com/go-ini/ini"
	"golang.org/x/crypto/bcrypt"
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

func GenPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func VerifyPassword(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
