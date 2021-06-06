package main

import (
	"encoding/json"
	"fmt"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/util"
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

	//err := db.Db.AutoMigrate(user.User{})
	//fmt.Println(err)

	//rand.Seed(time.Now().UnixNano())
	//var uid uint = 12
	//fmt.Println(strconv.FormatUint(uint64(uid), 10))
	//s := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", rand.Int()))))
	//s += uids
	//fmt.Printf("%T", s[32:])
	s, _ := json.Marshal(util.Success.Msg(user.User{Id: 10}))
	fmt.Printf("%s\n", s)
}
