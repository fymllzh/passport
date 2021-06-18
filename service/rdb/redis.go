package rdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/wuzehv/passport/util"
	"strconv"
	"time"
)

var Rdb *redis.Client

func init() {
	port, err := strconv.Atoi(util.ENV("redis", "db"))
	if err != nil {
		fmt.Printf("rdb db num config error '%v', use 0\n", err)
		port = 0
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     util.ENV("redis", "host") + ":" + util.ENV("redis", "port"),
		Password: util.ENV("redis", "passwd"),
		DB:       port,
	})

	var ctx = context.Background()
	_, err = Rdb.Ping(ctx).Result()
	// redis初始化错误
	if err != nil {
		panic(err)
	}
}

func SetJson(k string, v interface{}, expiration time.Duration) {
	var ctx = context.Background()
	str, _ := json.Marshal(v)
	Rdb.Set(ctx, k, str, expiration)
}

func GetJson(k string, v interface{}) bool {
	var ctx = context.Background()
	cache, err := Rdb.Get(ctx, k).Result()
	if err != nil {
		fmt.Println()
		return false
	}

	if cache != "" {
		json.Unmarshal([]byte(cache), &v)
		return true
	}

	return false
}
