package redis

import (
	"ecode/config"
	"fmt"

	"github.com/go-redis/redis"
)

// DB db
var DB *redis.Client

func init() {
	var config = config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.IP + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("redis 连接失败")
	}
	DB = client
	fmt.Println("redis 连接成功")
}
