package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type NotifController interface {
	GetAllNotif(ctx *gin.Context)
}

type notifController struct {
	notifService services.NotifService
	jwtService   services.JWTService
}

func NewNotifController(notifService services.NotifService, jwtService services.JWTService) NotifController {
	return &notifController{
		notifService: notifService,
		jwtService:   jwtService,
	}
}

func (c *notifController) GetAllNotif(ctx *gin.Context) {
	page, _ := strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userId := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userId, 10, 64)
	res, err := c.notifService.GetAllNotif(convertedUserID, page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	if res == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "menampilkan data",
			"data":    []string{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "menampilkan data",
		"data":    &res,
	})
}

func (c *notifController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
