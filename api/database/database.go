package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("PASSWORD"),
		DB:       dbNo,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}

	return rdb
}
