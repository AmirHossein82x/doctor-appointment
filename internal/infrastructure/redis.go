package infrastructure

import (
	"strconv"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		config := config.LoadConfig()
		RedisDb, err := strconv.Atoi(config.REDIS_DB)
		if err != nil {
			panic(err)
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.REDIS_ADDR,
			DB:   RedisDb,
		})
	}
	return redisClient
}
