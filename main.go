package main

import (
	"github.com/wuzehv/passport/router"
	"github.com/wuzehv/passport/service/db"
	"github.com/wuzehv/passport/util"
)

func main() {
	defer db.Db.Close()

	router := router.InitRouter()
	router.Run(util.ENV("product", "addr"))
}
