package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuzehv/passport/util"
	"log"
)

var Db *sql.DB

func init() {
	return
	user := util.ENV("db", "dbUser")
	passwd := util.ENV("db", "dbPasswd")

	host := util.ENV("db", "dbHost")
	port := util.ENV("db", "dbPort")

	db := util.ENV("db", "dbName")

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
