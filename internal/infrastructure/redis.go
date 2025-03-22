package infrastructure

import (
	"strconv"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	config := config.LoadConfig()
	RedisDb, err := strconv.Atoi(config.REDIS_DB)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{
		Addr: config.REDIS_ADDR,
		DB:   RedisDb,
	})
}
