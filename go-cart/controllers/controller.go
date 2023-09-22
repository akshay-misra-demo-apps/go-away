package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gocart.com/go-cart/utils"
)

type IController interface {
}

func ValidateToken(c *gin.Context) error {
	ok, err := utils.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(403, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
		return err
	}
	if !ok {
		c.JSON(403, gin.H{
			"message": fmt.Sprintf("authentication error: %v", err.Error()),
		})

		return err
	}

	return nil
}
