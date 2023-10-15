package cache

import (
	"github.com/redis/go-redis/v9"
	"go-boss/config"
)

var RedisCache *redis.Client

func InitRedis() {
	c := config.NewConfig()
	viper := c.Viper
	addr := viper.Get("redis.addr")
	password := viper.Get("redis.password")
	db := viper.Get("redis.db")
	RedisCache = redis.NewClient(&redis.Options{
		Addr:     addr.(string),
		Password: password.(string),
		DB:       db.(int),
	})
}
