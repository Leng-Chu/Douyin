package redis

import (
	"Douyin/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb1 *redis.Client
var rdb2 *redis.Client

func Init() {
	info := config.Conf.Redis

	rdb1 = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", info.Host, info.Port),
			Password: fmt.Sprintf("%s", info.Pass),
			DB:       1,
		})

	rdb2 = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", info.Host, info.Port),
			Password: fmt.Sprintf("%s", info.Pass),
			DB:       2,
		})
}
