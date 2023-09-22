package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gocart.com/go-cart/controllers"
	"gocart.com/go-cart/repositories"
	"gocart.com/go-cart/routes"
	"gocart.com/go-cart/services"
)

func init() {
	// check if this is good from DI perspective
	repositories.Init()
}

func App() {
	server := gin.Default()
	server.GET("/", home)
	productRepository := repositories.Get()["product"]
	productService := services.GetProductService(productRepository)
	productController := controllers.GetProductController(productService)
	routes.ProductRoutes(server, productController)

	userRepository := repositories.Get()["user"]
	userService := services.GetUserService(userRepository)
	userController := controllers.GetUserController(userService)
	routes.UserRoutes(server, userController)

	if err := server.Run(":4000"); err != nil {
		log.Fatalf("error while creating http server %v", err.Error())
	}
}

// consider using some DI like google.wire to manage dependencies.
func main() {
	App()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Welcome to go-cart!",
	})
}
