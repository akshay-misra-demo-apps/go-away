package repositories

import (
	"log"

	"gocart.com/go-cart/database"
)

const ConnectionURL = "mongodb+srv://userone:L2ORM855sl5XzunU@cluster0.9eqsnyg.mongodb.net"

// Repositories singleton map
type Repositories map[string]IRepository

var repositories Repositories

func Init() {
	if repositories == nil {
		repositories = make(map[string]IRepository)
	}

	mongo := database.Mongo{
		ConnectionString: ConnectionURL,
	}

	connection, err := mongo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	initProductRepo(connection)
}

func initProductRepo(connection database.DBConnection) {
	productRepository := ProductRepository{
		connection: connection,
	}
	repositories["product"] = productRepository
}

func Get() Repositories {
	if repositories == nil {
		log.Fatal("repositories not initialized!, call Init while initializing the application.")
	}

	return repositories
}
