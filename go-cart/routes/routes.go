package routes

import (
	"github.com/gin-gonic/gin"
	"gocart.com/go-cart/controllers"
)

func ProductRoutes(r *gin.Engine, c *controllers.ProductController) {
	// product catalog routes
	product := r.Group("/api/product")
	product.GET("/", c.GetProducts)
	product.GET("/:id", c.GetProduct)
	product.POST("/", c.CreateProduct)
	product.PATCH("/:id", c.PatchProduct)
	product.DELETE("/:id", c.DeleteProduct)
}
