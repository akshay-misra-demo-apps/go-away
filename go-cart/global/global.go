package global

import (
	"gocart.com/go-cart/services"
)

var (
	PRODUCT = services.Get()
	USER    = services.User{}
)
