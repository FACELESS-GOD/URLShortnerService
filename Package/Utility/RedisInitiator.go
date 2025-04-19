package Utility

import (
	Config "github.com/FACELESS-GOD/URLShortnerService.git/Helper/MetaData"
	"github.com/redis/go-redis/v9"
)

var RedisInstance *redis.Client

func InitaiteRedisInstance() {
	client := redis.NewClient(&redis.Options{
		Addr:     Config.RedisConnString,
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,
	})

	RedisInstance = client
}
