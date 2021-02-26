package main

import (
	"github.com/wuzehv/passport/router"
	"github.com/wuzehv/passport/util"
)

func main() {
	r := router.InitRouter()
	r.Run(util.ENV("", "addr"))
}
