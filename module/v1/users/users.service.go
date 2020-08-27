package users

import (
	"context"
	"log"
	"net/http"
	"strings"
	"ulum-go/config"
	"ulum-go/module/v1/users/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UsersService ...
type UsersService struct{}

var validate *validator.Validate

// Connection ...
func Connection() *mongo.Collection {
	db := config.Connect().Database("test").Collection("heroes")
	return db
}

// Register ...
func (UsersService) Register(c *gin.Context) {
	var heroes []*UsersSchema
	var json types.Register
	if err := c.ShouldBindJSON(&json); err != nil {
		split := strings.Split("asdasdas:asdasd", ":")
		println(split)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur, err := Connection().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero UsersSchema
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	for _, hero := range heroes {
		log.Println(hero.Username, hero.Password)
	}
	c.JSON(http.StatusOK, heroes)
}
