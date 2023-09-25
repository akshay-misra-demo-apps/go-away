package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Entity
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdateAt  string             `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Auth struct {
	Token        string `json:"username"`
	RefreshToken string `json:"username"`
	Id           string
	Username     string
	// Change to enum
	Status string
}
