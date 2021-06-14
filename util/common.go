package util

import (
	"crypto/md5"
	"fmt"
	"github.com/go-ini/ini"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sort"
	"strings"
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

type key []string

func (k key) Len() int {
	return len(k)
}

func (k key) Less(i, j int) bool {
	return k[j] > k[i]
}

func (k key) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func GenSign(m map[string]string, secret string) string {
	keyArr := make(key, len(m))
	i := 0
	for k, _ := range m {
		if k == Sign {
			continue
		}

		keyArr[i] = k
		i++
	}
	sort.Sort(keyArr)

	var b strings.Builder
	for _, v := range keyArr {
		b.WriteString(m[v])
	}
	b.WriteString(secret)

	return fmt.Sprintf("%X", md5.Sum([]byte(b.String())))
}
