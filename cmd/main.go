package main

import (
	// "context"
	// "fmt"
	// "log"

	"context"
	"fmt"

	"github.com/AmirHossein82x/doctor-appointment/internal/delivery"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/AmirHossein82x/doctor-appointment/internal/logger"

	"github.com/gin-gonic/gin"
	// "github.com/AmirHossein82x/doctor-appointment/internal/logger"
)

func main() {
	// cfg := config.LoadConfig()
	infrastructure.ConnectToDB()
	client := infrastructure.GetRedisClient()
	ctx := context.Background()
	// Ping the Redis server to check the connection
	client.Set(ctx, "key", "value", 10)
	val, err := client.Get(ctx, "key").
		Result()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("key", val)
	log := logger.SetUpLogger()
	log.Info("This is an info log")                           // Won't be sent to ES
	log.Warn("testttttttttttttttttttttttttttt a warning log") // Sent to Elasticsearch
	delivery.TestLog()

	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "working well")
	})
	app.Run(":8080")

}
