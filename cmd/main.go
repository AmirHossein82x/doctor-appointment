package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
)

func main() {
	// cfg := config.LoadConfig()
	infrastructure.ConnectToDB()
	client := infrastructure.NewRedisClient()
	ctx := context.Background()
	// Ping the Redis server to check the connection
	client.Set(ctx, "key", "value", 0)
	val, err := client.Get(ctx, "key").
		Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key", val)
}
