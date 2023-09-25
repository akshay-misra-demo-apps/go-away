package interfaces

import (
	"gocart.com/go-cart/models"
)

type IUser interface {
	Register(*models.User) (*models.SignupResponse, error)
	Authenticate(*models.Login) (*models.LoginResponse, error)
	Logout(string) error
	Get(string) (*models.User, error)
	Patch(*models.User) (*models.User, error)
	Delete(string) error
}
