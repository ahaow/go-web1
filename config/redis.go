package config

import (
	"github.com/go-redis/redis"
	"go-web1/global"
	"log"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", DB: 0, Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}
	global.RedisDB = RedisClient
}
