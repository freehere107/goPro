package utiles

import (
	"github.com/spf13/viper"
	"gopkg.in/redis.v4"
)

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
