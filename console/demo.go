package main

import (
	"encoding/json"
	"errors"
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
		token, err := r.Cookie("token")
		port := util.ENV("", "addr")
		domain := "http://" + util.ENV("", "domain")

		if err != nil {
			http.Redirect(w, r, domain+port+"/sso/index?callback=http://"+r.Host+"/login", http.StatusTemporaryRedirect)
			return
		} else {
			// 请求sso进行auth
			_, err := httpRequest("/sso/auth", token.Value)

			if err != nil {
				http.Redirect(w, r, domain+port+"/sso/index?callback=http://"+r.Host+"/login", http.StatusTemporaryRedirect)
				return
			}
		}
		handler(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	if _, ok := param["token"]; !ok {
		fmt.Fprintln(w, "system error")
		return
	}

	token := param["token"][0]

	_, err := httpRequest("/sso/session", token)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: token})

	http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
}

func _default(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "index page")
}

func httpRequest(url string, token string) (sso.User, error) {
	port := util.ENV("", "addr")
	domain := "http://" + util.ENV("", "domain")
	res, err := http.Get(domain + port + url + "?token=" + token)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var d sso.Response
	if err = json.Unmarshal(str, &d); err != nil {
		log.Fatalln(err)
	}

	if !d.Success {
		return d.Data, errors.New(d.Message)
	}

	return d.Data, nil
}
