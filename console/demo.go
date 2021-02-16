package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wuzehv/passport/app/sso"
	"github.com/wuzehv/passport/util"
	"io/ioutil"
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
		_, err := r.Cookie("name")
		port := util.ENV("", "addr")

		if err != nil {
			http.Redirect(w, r, "http://sso.com"+port+"/sso/index?callback=http://"+r.Host+"/login", http.StatusMovedPermanently)
			return
		} else {
			// 请求sso进行auth
			var d sso.Response
			httpRequest("/sso/auth", "token_str1", &d)

			if !d.Success {
				http.Redirect(w, r, "http://sso.com"+port+"/sso/index?callback=http://"+r.Host+"/login", http.StatusMovedPermanently)
				return
			}
		}
		handler(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	token := param["token"][0]

	var d sso.Response
	httpRequest("/sso/session", token, &d)

	if !d.Success {
		fmt.Fprintln(w, "logout fail")
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "name", Value: d.Data.Name})
	fmt.Fprintln(w, d.Data.Name)
}

func _default(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "default")
}

func httpRequest(url string, token string, value interface{}) {
	port := util.ENV("", "addr")
	res, err := http.Get("http://sso.com" + port + url + "?token=" + token)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(str, &value); err != nil {
		log.Fatalln(err)
	}
}
