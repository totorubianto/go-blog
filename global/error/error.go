package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandling ...
func ErrorHandling(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "error": message})
	return
}
