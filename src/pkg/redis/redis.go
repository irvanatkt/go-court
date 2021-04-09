package redis

import (
	"context"
	"fmt"
	"log"

	config "court.com/src/pkg/utils/conf"
	"github.com/go-redis/redis/v8"
)

func Init(conf config.Redis) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		DB:   conf.DB,
	})
	ping, err := cli.Ping(context.Background()).Result()
	log.Printf("Pinging to redis %s, err: %v", ping, err)
	return cli
}
