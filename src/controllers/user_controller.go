package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type UserController interface {
	CreatePassword(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) CreatePassword(ctx *gin.Context) {
	userCreatePasswordRequest := request.UserCreatePasswordRequest{}
	err := ctx.ShouldBind(&userCreatePasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err = c.userService.CreatedPassword(userCreatePasswordRequest, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Password berhasil dibuat",
	})
}

func (c *userController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
