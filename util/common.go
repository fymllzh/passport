package util

import (
	"crypto/md5"
	"fmt"
	"github.com/go-ini/ini"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
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

func GenToken() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", rand.Int()))))
}
