package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func APIKey() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("X_API_KEY")
	return func(ctx *gin.Context) {
		RequestHeaderName := ctx.Request.Header.Get("x-api-key")
		if RequestHeaderName != apiKey {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "X-API-KEY ERROR",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
