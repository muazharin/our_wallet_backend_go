package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) <= 0 {
		c.Next()
		return
	}
	for _, err := range c.Errors {
		log.Printf("Error -> %+v\n", err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  false,
		"message": "StatusInternalServerError",
	})
}
