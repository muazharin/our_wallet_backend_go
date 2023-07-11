package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

func AuthorizeJWT(jwtService services.JWTService, authRepo repositories.AuthRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Token tidak ditemukan",
			})
			ctx.Abort()
			return
		}
		authHeader = strings.Split(authHeader, "Bearer ")[1]
		token, e := jwtService.ValidateToken(authHeader)
		if e != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Token tidak valid",
			})
			ctx.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			user_phone := fmt.Sprintf("%v", claims["user_phone"])
			res, _, _ := authRepo.CheckAccount(user_phone, user_phone, user_phone)
			if res < 1 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status":  false,
					"message": "User tidak ditemukan",
				})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Token tidak valid",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
