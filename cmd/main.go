package main

import (
	"github.com/AmirHossein82x/doctor-appointment/internal/delivery"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.ConnectToDB()

	app := gin.Default()
	delivery.SetUpUserRoutes(&app.RouterGroup)

	app.Run(":8080")

}
