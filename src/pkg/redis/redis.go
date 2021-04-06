package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func Init() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	cmd := cli.Ping(context.Background())
	log.Println(cmd.Result())
	return cli
}
