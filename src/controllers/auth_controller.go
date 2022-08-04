package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type AuthController interface {
	CheckPhone(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) CheckPhone(ctx *gin.Context) {
	var chekPhoneRequest request.AuthCheckPhoneRequest
	err := ctx.ShouldBind(&chekPhoneRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	res, err := c.authService.CheckPhone(chekPhoneRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if res > 0 {
		ctx.JSON(http.StatusCreated, gin.H{
			"status":  false,
			"message": "Nomor Anda sudah terdaftar",
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Nomor Anda boleh digunakan",
		})

	}
	return
}

func (c *authController) SignUp(ctx *gin.Context) {
	var authSignUpRequest request.AuthSignUpRequest
	err := ctx.ShouldBind(&authSignUpRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	res, err := c.authService.SignUp(authSignUpRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	generatedToken := c.jwtService.GenerateToken(res)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"token":   &generatedToken,
	})
	return
}
