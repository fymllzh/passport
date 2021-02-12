package main

import (
	"flag"
	"fmt"
	"github.com/wuzehv/passport/util"
	"log"
	"net/http"
)

var (
	h    bool
	addr string
)

func main() {
	flag.BoolVar(&h, "h", false, "usage")
	flag.StringVar(&addr, "addr", "127.0.0.1:8081", "listen address")

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	fmt.Println("listen:", addr)

	http.HandleFunc("/", wrapHandler(_default))
	http.HandleFunc("/login", login)

	log.Fatalln(http.ListenAndServe(addr, nil))
}

func wrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("wzh")
		if err != nil {
			port := util.ENV("", "addr")
			http.Redirect(w, r, "http://sso.com" + port + "/sso/index?callback=http://" + r.Host + "/login", http.StatusMovedPermanently)
			return
		} else {
			// 请求sso进行auth
		}
		handler(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "login")
}

func _default(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "default")
}
