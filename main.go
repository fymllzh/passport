package main

import (
	"github.com/wuzehv/passport/router"
	"github.com/wuzehv/passport/util/config"
	"log"
)

func main() {
	r := router.InitRouter()
	if err := r.Run(config.App.Port); err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
}
