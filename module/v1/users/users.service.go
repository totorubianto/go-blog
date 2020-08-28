package users

import (
	"context"
	"log"
	"net/http"
	"ulum-go/global/error"
	"ulum-go/module/v1/users/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

// UsersService ...
type UsersService struct{}

var validate *validator.Validate

// Register ...
func (UsersService) Register(c *gin.Context) {
	var json types.Register
	if err := c.ShouldBindJSON(&json); err != nil {
		error.ErrorHandling(c, err.Error())
		return
	}
	ctx := context.Background()
	cur, err := Connection().InsertOne(ctx, bson.M{"username": json.Username, "password": json.Password, "fullName": json.Fullname, "email": json.Email})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	c.JSON(http.StatusOK, cur)
}

// Find ...
func (UsersService) Find(c *gin.Context) {
	var users []*UsersSchema
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
		users = append(users, &hero)
	}
	c.JSON(http.StatusOK, users)
}

// Login ...
func (UsersService) Login(c *gin.Context) {
	var json types.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		error.ErrorHandling(c, err.Error())
		return
	}
	var password string = json.Password
	println(password)
	ctx := context.Background()
	_ = Connection().FindOne(ctx, bson.M{"username": json.Username}).Decode(&json)
	if password != json.Password {
		error.ErrorHandling(c, "Password salah!!!!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "data": json})
}
