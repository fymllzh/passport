// 初始化数据库和表结构
package main

import (
	"github.com/wuzehv/passport/service/db"
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
	_, err := db.Db.Exec(passport)
	if err != nil {
		panic(err)
	}
	log.Println("create database done")
}
