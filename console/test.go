package main

import (
	"fmt"
	"github.com/wuzehv/passport/util"
)

func main() {
	hash := util.GenPassword("abc")
	fmt.Println(util.VerifyPassword(hash, "abc"))
}
