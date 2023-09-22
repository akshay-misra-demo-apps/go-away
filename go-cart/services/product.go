package services

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gocart.com/go-cart/interfaces"
	"gocart.com/go-cart/models"
	"gocart.com/go-cart/repositories"
)

type ProductService struct {
	repository repositories.IRepository
}

func GetProductService(repository repositories.IRepository) interfaces.IProduct {
	return &ProductService{
		repository: repository,
	}
}

func (p *ProductService) Create(product *models.Product) (*models.Product, error) {
	if json, err := json2.Marshal(product); err == nil {
		log.Printf("created new product: %v", string(json))
	}

	product.Id = primitive.NewObjectID()

	_, err := p.repository.GetCollection().InsertOne(context.TODO(), product)
	if err != nil {
		return nil, fmt.Errorf("error while inserting new product to database: %v", err)
	}

	return product, nil
}

func (p *ProductService) Get(id string) (*models.Product, error) {

	out := &models.Product{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := p.repository.GetCollection().FindOne(context.TODO(), bson.M{"_id": objectId})

	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	result.Decode(out)

	return out, nil
}

func (p *ProductService) GetAll() ([]*models.Product, error) {
	return []*models.Product{}, nil
}
func (p *ProductService) Search(regex string) ([]*models.Product, error) {
	return []*models.Product{}, nil
}

func (p *ProductService) Patch(product *models.Product) (*models.Product, error) {

	if json, err := json2.Marshal(product); err == nil {
		log.Printf("patched existing product: %v", string(json))
	}

	var updates map[string]interface{}
	json, err := json2.Marshal(product)
	if err != nil {
		return nil, err
	}

	err = json2.Unmarshal(json, &updates)
	if err != nil {
		return nil, err
	}

	patch := bson.M{"$set": updates}
	filter := bson.D{{"_id", product.Id}}

	_, err = p.repository.GetCollection().UpdateOne(context.TODO(), filter, patch)
	if err != nil {
		return nil, fmt.Errorf("error while patching existing product to database: %v", err)
	}

	return product, nil
}

func (p *ProductService) Delete(id string) error {
	return nil
}
