package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gocart.com/go-cart/database"
)

type ProductRepository struct {
	connection database.DBConnection
}

// GetCollection Call this everytime we want to interact with the mongo collection, like insert, read etc.
func (r ProductRepository) GetCollection() *mongo.Collection {
	// Get Products collection
	return r.connection.Client.
		Database("go-cart").
		Collection("Products")
}
