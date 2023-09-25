package interfaces

import (
	"gocart.com/go-cart/models"
)

type IUser interface {
	Register(*models.User) (*models.User, error)
	Authenticate(*models.AuthRequest) (*models.Auth, error)
	Logout(string) error
	Get(string) (*models.User, error)
	Patch(*models.User) (*models.User, error)
	Delete(string) error
}
