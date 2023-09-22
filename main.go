package main

import (
	"AnaShopService/config"
	"AnaShopService/controller"
	"AnaShopService/docs"
	"AnaShopService/repository"
	"AnaShopService/service"
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	db := config.GetDB()

	shopRepository := repository.NewShopRepository(db)
	shopService := service.NewShopService(shopRepository)
	shopController := controller.NewShopController(shopService)

	docs.SwaggerInfo.Title = "Ana Store - Environment: "
	docs.SwaggerInfo.Description = "API for shop service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/cart", shopController.CartRequest)

	// Start the server
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
