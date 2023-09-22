package routes

import (
	"github.com/gin-gonic/gin"
	"gocart.com/go-cart/controllers"
)

func UserRoutes(r *gin.Engine, c *controllers.UserController) {
	// user management routes
	product := r.Group("/api/user")
	product.POST("/register", c.Register)
	product.POST("/login", c.Login)
	product.POST("/logout", c.Logout)
	product.GET("/:id", c.GetUser)
	product.PATCH("/:id", c.PatchUser)
	product.DELETE("/:id", c.DeleteUser)
}

func ProductRoutes(r *gin.Engine, c *controllers.ProductController) {
	// product catalog management routes
	product := r.Group("/api/product")
	product.GET("/", c.GetProducts)
	product.GET("/:id", c.GetProduct)
	product.POST("/", c.CreateProduct)
	product.PATCH("/:id", c.PatchProduct)
	product.DELETE("/:id", c.DeleteProduct)
}
