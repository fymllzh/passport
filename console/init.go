// 初始化数据库和表结构
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuzehv/passport/model/base"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/session"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dropDb = `
DROP DATABASE IF EXISTS passport
`

var passportDb = `
CREATE DATABASE IF NOT EXISTS passport DEFAULT CHARACTER SET utf8
`

func main() {
	u := util.ENV("db", "user")
	passwd := util.ENV("db", "passwd")

	host := util.ENV("db", "host")
	port := util.ENV("db", "port")

	dbName := util.ENV("db", "name")

	dsn := u + ":" + passwd + "@tcp(" + host + ":" + port + ")/?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.Exec(dropDb)
	log.Println("drop database done")

	db.Exec(passportDb)
	log.Println("create database done")

	db.Exec(passportDb)
	log.Println("create database done")

	dsn = u + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(user.User{})
	if err != nil {
		panic(err)
	}

	log.Println("create users table done")

	err = db.AutoMigrate(session.Session{})
	if err != nil {
		panic(err)
	}

	log.Println("create tokens table done")

	err = db.AutoMigrate(client.Client{})
	if err != nil {
		panic(err)
	}

	log.Println("create clients table done")

	db.Create(&client.Client{Domain: "client.one.com", Callback: "http://client.one.com:8081/callback", Secret: "123456", Status: base.StatusNormal})
	db.Create(&client.Client{Domain: "client.two.com", Callback: "http://client.one.com:8082/callback", Secret: "123456", Status: base.StatusNormal})

	log.Println("initialize client done")

	u = "admin@gmail.com"
	up := "admin"
	p := util.GenPassword(up)
	db.Create(&user.User{Email: u, Password: p})
	log.Printf("initialize user: %s, password: %s\n", u, up)
}
