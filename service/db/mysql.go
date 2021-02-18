package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuzehv/passport/util"
	"log"
)

var Db *sql.DB

func init() {
	user := util.ENV("db", "user")
	passwd := util.ENV("db", "passwd")

	host := util.ENV("db", "host")
	port := util.ENV("db", "port")

	db := util.ENV("db", "name")

	var err error
	Db, err = sql.Open("mysql", user+":"+passwd+"@tcp("+host+":"+port+")/"+db+"?parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}

	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)

	if err := Db.Ping(); err != nil {
		log.Fatalln(err)
	}
}
