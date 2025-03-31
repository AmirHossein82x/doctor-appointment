package main

import (
	_ "github.com/AmirHossein82x/doctor-appointment/docs" // Import generated Swagger docs
	"github.com/AmirHossein82x/doctor-appointment/internal/delivery"
	"github.com/AmirHossein82x/doctor-appointment/internal/infrastructure"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	infrastructure.ConnectToDB()

	app := gin.Default()
	// Swagger route
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	delivery.SetUpOtpRoutes(&app.RouterGroup)
	delivery.SetUpUserRoutes(&app.RouterGroup)
	delivery.SetUpAdminRoutes(&app.RouterGroup)

	app.Run(":8080")

}
