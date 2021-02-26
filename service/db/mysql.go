package db

import (
	"github.com/wuzehv/passport/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	user := util.ENV("db", "user")
	passwd := util.ENV("db", "passwd")

	host := util.ENV("db", "host")
	port := util.ENV("db", "port")

	db := util.ENV("db", "name")

	var err error
	dsn := user + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
}
