package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"

	// "github.com/muazharin/our_wallet_backend_go/src/middlewares"
	"github.com/muazharin/our_wallet_backend_go/src/services"
)

type AuthController interface {
	CheckPhone(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	SignOut(ctx *gin.Context)
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
	var authResponse response.AuthResponse
	err := ctx.ShouldBind(&chekPhoneRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	count, res, err := c.authService.CheckPhone(chekPhoneRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	authResponse.Token = c.jwtService.GenerateToken(res)
	authResponse.UserStatus = res.UserStatus
	if count > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Nomor Anda sudah terdaftar",
			"data":    &authResponse,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Nomor Anda boleh digunakan",
		})

	}
}

func (c *authController) SignUp(ctx *gin.Context) {
	var authSignUpRequest request.AuthSignUpRequest
	var authResponse response.AuthResponse
	err := ctx.ShouldBind(&authSignUpRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
	authResponse.Token = c.jwtService.GenerateToken(res)
	authResponse.UserStatus = res.UserStatus

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Selamat! Anda telah terdaftar",
		"data":    &authResponse,
	})
}

func (c *authController) SignIn(ctx *gin.Context) {
	var authSignInRequest request.AuthSignInRequest
	var authResponse response.AuthResponse
	err := ctx.ShouldBind(&authSignInRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	res, err := c.authService.SignIn(authSignInRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	authResponse.Token = c.jwtService.GenerateToken(res)
	authResponse.UserStatus = res.UserStatus
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Selamat! Anda berhasil login",
		"data":    &authResponse,
	})
}

func (c *authController) SignOut(ctx *gin.Context) {
	var authSignOutRequest request.AuthSignOutRequest
	err := ctx.ShouldBind(&authSignOutRequest)
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
	err = c.authService.SignOut(authSignOutRequest, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Logout berhasil!",
	})
}

func (c *authController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
