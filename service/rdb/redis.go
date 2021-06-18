package rdb

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/wuzehv/passport/util/config"
	"log"
	"time"
)

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Passwd,
		DB:       config.Redis.DbNum,
	})

	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()
	// redis初始化错误
	if err != nil {
		log.Fatalf("redis init error: %v\n", err)
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
	if err != nil && err != redis.Nil {
		log.Printf("redis error: %v\n", err)
		return false
	}

	if cache != "" {
		json.Unmarshal([]byte(cache), &v)
		return true
	}

	return false
}
