package main

import (
	// "context"
	// "fmt"
	// "log"

	// "github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"

	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/delivery"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"
	// "github.com/AmirHossein82x/doctor-appointment/internal/logger"
)

func main() {
	// cfg := config.LoadConfig()
	// infrastructure.ConnectToDB()
	// client := infrastructure.NewRedisClient()
	// ctx := context.Background()
	// // Ping the Redis server to check the connection
	// client.Set(ctx, "key", "value", 0)
	// val, err := client.Get(ctx, "key").
	// 	Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("key", val)





	log := logger.SetUpLogger()
	log.Info("This is an info log")                            // Won't be sent to ES
	log.Warn("is a warning log")                              // Sent to Elasticsearch
	delivery.TestLog()
	
	time.Sleep(time.Second * 2)
}
