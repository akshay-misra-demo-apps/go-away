package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Entity
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name,required"`
	Username  string             `json:"username" bson:"username,required"`
	Password  string             `json:"password" bson:"password,required"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdateAt  string             `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
}

type Auth struct {
	Token    string
	Id       string
	Username string
	// Change to enum
	Status string
}
