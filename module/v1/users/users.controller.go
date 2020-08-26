package users

import (
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(r *gin.RouterGroup) {
	users := new(UsersService)
	router := r.Group("/users")
	{
		router.POST("/register", users.Register)
	}
}
