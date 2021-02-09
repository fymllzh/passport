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
	user := util.ENV("db", "dbUser")
	passwd := util.ENV("db", "dbPasswd")

	host := util.ENV("db", "dbHost")
	port := util.ENV("db", "dbPort")

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