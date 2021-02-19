// 初始化数据库和表结构
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuzehv/passport/util"
	"log"
)

var passportDb = `
CREATE DATABASE IF NOT EXISTS passport DEFAULT CHARACTER SET utf8
`

var userTable = `
CREATE TABLE IF NOT EXISTS passport.user
(
  id       int AUTO_INCREMENT
    PRIMARY KEY,
  email    varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  CONSTRAINT email
    UNIQUE (email)
) CHARACTER SET utf8
`

var tokenTable = `
CREATE TABLE IF NOT EXISTS passport.token
(
  id      int AUTO_INCREMENT
    PRIMARY KEY,
  user_id int                                NOT NULL,
  token   varchar(255)                       NOT NULL,
  created datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT token
    UNIQUE (token)
) CHARACTER SET utf8
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

	_, err = db.Exec(passportDb)
	if err != nil {
		panic(err)
	}

	log.Println("create database done")

	_, err = db.Exec(userTable)
	if err != nil {
		panic(err)
	}

	log.Println("create user table done")

	_, err = db.Exec(tokenTable)
	if err != nil {
		panic(err)
	}

	log.Println("create token table done")

	u := "admin"
	p := util.GenPassword("admin")
	usql := fmt.Sprintf("insert into passport.user values (null, '%s', '%s')", u, p)
	_, err = db.Exec(usql)
	if err != nil {
		panic(err)
	}

	log.Println("initialize user: admin, password: admin")
}
