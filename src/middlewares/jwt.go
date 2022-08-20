package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "No token found",
			})
			return
		}
		authHeader = strings.Split(authHeader, "Bearer ")[1]
		token, _ := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Token is not valid",
			})
		}
	}
}
