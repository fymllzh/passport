package main

import (
	"fmt"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/util"
)

func main() {
	hash := util.GenPassword("abc")
	fmt.Println(util.VerifyPassword(hash, "abc"))

	s, err := user.FindByEmail("admin")
	fmt.Println(s, err)
}
