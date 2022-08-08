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

type WalletController interface {
	GetAllWallet(ctx *gin.Context)
	CreateWallet(ctx *gin.Context)
}

type walletController struct {
	walletService services.WalletService
	jwtService    services.JWTService
}

func NewWalletController(walletService services.WalletService, jwtService services.JWTService) WalletController {
	return &walletController{
		walletService: walletService,
		jwtService:    jwtService,
	}
}

func (c *walletController) CreateWallet(ctx *gin.Context) {
	var createwallet request.WalletCreateReq
	err := ctx.ShouldBind(&createwallet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err = c.walletService.CreateWallet(createwallet, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Wallet berhasil dibuat",
	})
}

func (c *walletController) GetAllWallet(ctx *gin.Context) {
	page, _ := strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userId := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userId, 10, 64)
	res, err := c.walletService.GetAllWallet(convertedUserID, page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data tidak ditemukan",
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

func (c *walletController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
