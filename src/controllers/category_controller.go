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

type CategoryController interface {
	AddCategory(ctx *gin.Context)
	GetAllCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
	jwtService      services.JWTService
}

func NewCategoryController(categoryService services.CategoryService, jwtService services.JWTService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
		jwtService:      jwtService,
	}
}

func (c *categoryController) AddCategory(ctx *gin.Context) {
	categoryAddRequest := request.CategoryAddRequest{}
	err := ctx.ShouldBind(&categoryAddRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userId := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userId, 10, 64)
	err = c.categoryService.AddCategory(categoryAddRequest, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Kategori berhasil ditambahkan",
	})
}

func (c *categoryController) GetAllCategory(ctx *gin.Context) {
	var categoryGetAllRequest request.CategoryGetAllRequest
	categoryGetAllRequest.CategoryType = ctx.Request.URL.Query().Get("type")
	categoryGetAllRequest.CategoryWalletId = ctx.Request.URL.Query().Get("walletId")
	categoryGetAllRequest.CategoryPage, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)

	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userId := c.getUserIDByToken(authHeader)
	categoryGetAllRequest.CategoryUserId = userId

	res, err := c.categoryService.GetAllCategory(categoryGetAllRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
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

func (c *categoryController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
