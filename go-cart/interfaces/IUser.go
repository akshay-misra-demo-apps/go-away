package interfaces

import (
	"gocart.com/go-cart/models"
)

type IUser interface {
	Register(*models.User) (*models.User, error)
	Authenticate(*models.User) (*models.Auth, error)
	Logout(string) error
	GetUser(string) (*models.User, error)
	PatchUser(*models.User) (*models.User, error)
	DeleteUser(string) error
}
