package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type OWController interface {
	GetOwUser(ctx *gin.Context)
	GetForMember(ctx *gin.Context)
}

type owController struct {
	owService  services.OWService
	jwtService services.JWTService
}

func NewOWController(owService services.OWService, jwtService services.JWTService) OWController {
	return &owController{
		owService:  owService,
		jwtService: jwtService,
	}
}

func (c *owController) GetOwUser(ctx *gin.Context) {
	var owGetUserReq request.OwGetUserReq
	owGetUserReq.Page, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	owGetUserReq.WalletId, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("wallet_id"), 10, 64)
	err := ctx.ShouldBind(&owGetUserReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	res, err := c.owService.GetOwUser(owGetUserReq)
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

func (c *owController) GetForMember(ctx *gin.Context) {
	var owGetUserReq request.OwGetUserReq
	owGetUserReq.Page, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	owGetUserReq.WalletId, _ = strconv.ParseInt(ctx.Request.URL.Query().Get("wallet_id"), 10, 64)
	owGetUserReq.Keyword = ctx.Request.URL.Query().Get("keyword")
	err := ctx.ShouldBind(&owGetUserReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	res, err := c.owService.GetForMember(owGetUserReq)
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
