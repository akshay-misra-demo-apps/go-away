package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gocart.com/go-cart/interfaces"
	"gocart.com/go-cart/models"
)

type ProductController struct {
	Service interfaces.IProduct
}

func GetProductController(service interfaces.IProduct) *ProductController {
	return &ProductController{
		Service: service,
	}
}

func (p ProductController) GetProducts(c *gin.Context) {
	fmt.Println("get all products")

	products, err := p.Service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}

	c.JSON(200, gin.H{
		"data": products,
	})
}

func (p ProductController) GetProduct(c *gin.Context) {
	fmt.Println("get product by id", c.Param("id"))

	if c.Param("id") == "" {
		c.JSON(400, gin.H{
			"message": "product 'id' is missing in request!",
		})
	}

	product, err := p.Service.Get(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

func (p ProductController) CreateProduct(c *gin.Context) {
	fmt.Println("create new product")
	var product models.Product
	c.Bind(&product)

	product.Id = primitive.NewObjectID()

	created, err := p.Service.Create(&product)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}

	c.JSON(201, gin.H{
		"data": created,
	})
}

func (p ProductController) PatchProduct(c *gin.Context) {
	fmt.Println("patch existing product", c.Param("id"))

	if c.Param("id") == "" {
		c.JSON(400, gin.H{
			"message": "product 'id' is missing in request!",
		})
	}

	var product models.Product
	c.Bind(&product)

	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}
	product.Id = objectId

	updated, err := p.Service.Patch(&product)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}

	c.JSON(200, gin.H{
		"data": updated,
	})
}

func (p ProductController) DeleteProduct(c *gin.Context) {
	fmt.Println("delete product with id: ", c.Param("id"))

	if c.Param("id") == "" {
		c.JSON(400, gin.H{
			"message": "product 'id' is missing in request!",
		})
	}

	err := p.Service.Delete(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}

	c.JSON(200, gin.H{
		"message": "product has been deleted successfully!",
	})
}
