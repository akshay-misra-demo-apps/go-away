package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var ()

func Register(c *gin.Context) {
	fmt.Println("register route")
}

func Login(c *gin.Context) {
	fmt.Println("login route")
}

func Logout(c *gin.Context) {
	fmt.Println("logout route")
}

func Profile(c *gin.Context) {
	fmt.Println("profile with id: ", c.Param("id"))
}
