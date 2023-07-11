package utils

import (
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     V.GetString("redis.addr"),
		Password: V.GetString("redis.pass"),
		DB:       V.GetInt("redis.db"),
	})
	RDB = rdb
}
