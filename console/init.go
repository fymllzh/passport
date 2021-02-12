// 初始化数据库和表结构
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuzehv/passport/util"
	"log"
)

var passport = `
CREATE DATABASE IF NOT EXISTS passport DEFAULT CHARACTER SET utf8
`

var user = `
`

var token = `
`

func main() {
	user := util.ENV("db", "user")
	passwd := util.ENV("db", "passwd")

	host := util.ENV("db", "host")
	port := util.ENV("db", "port")

	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+host+":"+port+")/?parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = db.Exec(passport)
	if err != nil {
		panic(err)
	}

	log.Println("create database done")
}
