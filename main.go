package main

import (
	"context"
	"log"
	"ulum-go/config"
	"ulum-go/module/v1/users"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	c := config.Connect()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		users.Router(v1)
	}
	router.Run()
}
