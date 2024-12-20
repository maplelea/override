package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"override/config"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Viper.GetString("redis.addr"),
		Password: config.Viper.GetString("redis.password"),
		DB:       config.Viper.GetInt("redis.db"),
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(any(err))
	}
}
