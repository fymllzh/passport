package main

import (
	"fmt"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
)

func main() {
	//hash := util.GenPassword("abc")
	//fmt.Println(util.VerifyPassword(hash, "abc"))
	//
	//var u user.User
	//db.Db.Table(user.Table).Where("email = ?", "admin").First(&u)
	//fmt.Println(u.Password)
	//
	//s := fmt.Sprintf("%d", time.Now().UnixNano())
	//fmt.Printf("%T", s)

	err := db.Db.AutoMigrate(user.User{})
	fmt.Println(err)
}
