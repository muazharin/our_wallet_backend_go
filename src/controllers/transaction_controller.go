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

type TransController interface {
	CreateTransaction(ctx *gin.Context)
}

type transController struct {
	transService services.TransService
	jwtService   services.JWTService
}

func NewTransactionController(transService services.TransService, jwtService services.JWTService) TransController {
	return &transController{
		transService: transService,
		jwtService:   jwtService,
	}
}

func (c *transController) CreateTransaction(ctx *gin.Context) {
	var transCreateReq request.TransCreateReq
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err := ctx.ShouldBind(&transCreateReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	// mengecek extension file
	for _, v := range transCreateReq.TransFile {
		switch v.Header["Content-Type"][0] {
		case "image/png":
		case "image/jpg":
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "File tidak sesuai",
			})
			return
		}
	}

	res, err := c.transService.CreateTransaction(transCreateReq, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println(len(transCreateReq.TransFile))
	for i, v := range transCreateReq.TransFile {
		path := fmt.Sprintf("src/images/trFiles/%s", res[i].TfFile)
		fmt.Println(i)
		fmt.Println(res[i].TfFile)
		fmt.Println(path)
		if err := ctx.SaveUploadedFile(v, path); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Gagal upload gambar :" + err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
	})
}

func (c *transController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
