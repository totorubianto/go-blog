package users

import (
	"ulum-go/config"

	"go.mongodb.org/mongo-driver/mongo"
)

// Connection ...
func Connection() *mongo.Collection {
	db := config.Connect().Database("test").Collection("heroes")
	return db
}
