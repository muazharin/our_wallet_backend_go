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
	GetAllTransByWalletId(ctx *gin.Context)
	GetAllTransByUserId(ctx *gin.Context)
	GetTransById(ctx *gin.Context)
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
			"message": "Lengkapi form dengan baik dan benar",
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

	// menyimpan gambar
	for i, v := range transCreateReq.TransFile {
		path := fmt.Sprintf("src/images/trFiles/%s", res[i].TfFile)
		if err := ctx.SaveUploadedFile(v, path); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Gagal upload gambar :" + err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Berhasil melakukan transaksi",
	})
}

func (c *transController) GetAllTransByWalletId(ctx *gin.Context) {
	var transWallet request.TransByWalletIdReq
	transWallet.TransWalletId, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("wallet_id"), 10, 64)
	transWallet.Page, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	err := ctx.ShouldBind(&transWallet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	res, err := c.transService.GetAllTransByWalletId(transWallet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data berhasil ditampilkan",
			"data":    []string{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data berhasil ditampilkan",
		"data":    &res,
	})
}

func (c *transController) GetAllTransByUserId(ctx *gin.Context) {
	var transUser request.TransByUserIdReq
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	transUser.TransUserId, _ = strconv.ParseInt(userID, 10, 64)
	transUser.Page, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	err := ctx.ShouldBind(&transUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	res, err := c.transService.GetAllTransByUserId(transUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data berhasil ditampilkan",
			"data":    []string{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data berhasil ditampilkan",
		"data":    &res,
	})
}

func (c *transController) GetTransById(ctx *gin.Context) {
	var trans request.TransByIdReq
	trans.TransId, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("trans_id"), 10, 64)
	err := ctx.ShouldBind(&trans)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	res, err := c.transService.GetTransById(trans)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data berhasil ditampilkan",
		"data":    &res,
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
