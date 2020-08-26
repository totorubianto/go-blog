package users

import (
	"context"
	"log"
	"net/http"
	"ulum-go/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// UsersService ...
type UsersService struct{}

// Register ...
func (UsersService) Register(c *gin.Context) {
	var heroes []*UsersSchema
	db := config.Connect().Database("test").Collection("heroes")
	cur, err := db.Find(context.TODO(), bson.M{})
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
