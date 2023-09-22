package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gocart.com/go-cart/database"
)

type UserRepository struct {
	connection database.DBConnection
}

// GetCollection Call this everytime we want to interact with the mongo collection, like insert, read etc.
func (r UserRepository) GetCollection() *mongo.Collection {
	// Get Users collection
	return r.connection.Client.
		Database("go-cart").
		Collection("Users")
}

func initUserRepo(connection database.DBConnection) {
	userRepository := UserRepository{
		connection: connection,
	}
	repositories["user"] = userRepository
}
