package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type CategoryController interface {
	AddCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
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
	err = c.categoryService.AddCategory(categoryAddRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Kategori berhasil ditambahkan",
	})
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	err := ctx.ShouldBind(&categoryUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	err = c.categoryService.UpdateCategory(categoryUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Kategori berhasil diperbaharui",
	})
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	categoryDeleteRequest := request.CategoryDeleteRequest{}
	err := ctx.ShouldBind(&categoryDeleteRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	err = c.categoryService.DeleteCategory(categoryDeleteRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Kategori berhasil dihapus",
	})
}

func (c *categoryController) GetAllCategory(ctx *gin.Context) {
	var categoryGetAllRequest request.CategoryGetAllRequest
	categoryGetAllRequest.CategoryType = ctx.Request.URL.Query().Get("type")
	categoryGetAllRequest.CategoryWalletId = ctx.Request.URL.Query().Get("walletId")
	categoryGetAllRequest.CategoryPage, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)

	res, err := c.categoryService.GetAllCategory(categoryGetAllRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
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
