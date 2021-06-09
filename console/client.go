package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/wuzehv/passport/util"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
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
	http.HandleFunc("/callback", callback)

	log.Fatalln(http.ListenAndServe(addr, nil))
}

var domain string

var username string

func wrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		domain = url2.QueryEscape(r.Host)

		jump := url2.QueryEscape("http://" + r.Host + r.RequestURI)
		url := "http://" + util.ENV("", "domain") + util.ENV("", "addr") + "/sso/index?domain=" + domain + "&jump=" + jump

		if err != nil {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		} else {
			// 获取用户信息
			u, err := httpRequest("/svc/userinfo", token.Value)
			if err != nil {
				http.Redirect(w, r, url, http.StatusTemporaryRedirect)
				return
			}

			for v, item := range u.(map[string]interface{}) {
				if v == "email" {
					username = item.(string)
					break
				}
			}
		}
		handler(w, r)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	if _, ok := param["token"]; !ok {
		fmt.Fprintln(w, "system error")
		return
	}

	token := param["token"][0]

	_, err := httpRequest("/svc/session", token)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: token, Path: "/", Domain: domain})

	http.Redirect(w, r, param["jump"][0], http.StatusTemporaryRedirect)
}

func _default(w http.ResponseWriter, r *http.Request) {
	d, _ := url2.QueryUnescape(domain)
	fmt.Fprintln(w, "<h1>登录成功, 客户端: "+ d + "</h1>", "<h2>当前用户: "+username+"</h2>")
}

func httpRequest(url string, token string) (interface{}, error) {
	port := util.ENV("", "addr")
	ssoDomain := "http://" + util.ENV("", "domain")

	ssoUrl := ssoDomain + port + url + "?token=" + token + "&domain=" + domain

	log.Println(ssoUrl)

	res, err := http.Post(ssoUrl, "", nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("bb", err)
	}

	log.Printf("get response: %s\n", str)

	var d util.Response
	if err = json.Unmarshal(str, &d); err != nil {
		log.Fatalln(d, err)
	}

	if d.Code != 0 {
		return nil, errors.New(d.Message)
	}

	return d.Data, nil
}