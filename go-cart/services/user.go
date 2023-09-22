package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gocart.com/go-cart/interfaces"
	"gocart.com/go-cart/models"
	"gocart.com/go-cart/repositories"
	"gocart.com/go-cart/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repositories.IRepository
}

func GetUserService(repository repositories.IRepository) interfaces.IUser {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Register(user *models.User) (*models.User, error) {

	user.Id = primitive.NewObjectID()

	hashed, err := utils.Encrypt(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while inserting new user to database: %v", err)
	}

	user.Password = hashed
	user.Status = "active"
	user.CreatedAt = time.Now().Local().String()

	_, err = u.repository.GetCollection().InsertOne(context.TODO(), user)
	if err != nil {
		return nil, fmt.Errorf("error while inserting new user to database: %v", err)
	}

	return user, nil

}

func (u *UserService) Authenticate(user *models.AuthRequest) (auth *models.Auth, err error) {
	out := &models.User{}

	result := u.repository.GetCollection().FindOne(context.TODO(),
		bson.M{
			"username": user.Username,
		})

	if result.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user %v not found", user.Username)
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	result.Decode(out)

	if utils.Compare(user.Password, out.Password) {
		token, err := utils.GenerateToken(out)
		if err != nil {
			return nil, fmt.Errorf("error while generating JWT token for user: %v", user.Username)
		}
		auth = &models.Auth{
			Id:       out.Id.String(),
			Username: out.Username,
			Status:   out.Status,
			Token:    token,
		}
	} else {
		return nil, fmt.Errorf("wrong password for user: %v", user.Username)
	}

	return
}

func (u *UserService) Logout(id string) error {
	return nil
}
func (u *UserService) GetUser(string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) PatchUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (u *UserService) DeleteUser(id string) error {
	return nil
}
