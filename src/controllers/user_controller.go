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
	GetUserProfile(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
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
	err = c.userService.CreatedPassword(userCreatePasswordRequest, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Password berhasil dibuat",
	})
}

func (c *userController) GetUserProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	res, err := c.userService.GetUserProfile(convertedUserID)
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

func (c *userController) UpdatePhoto(ctx *gin.Context) {
	var userPhotoReq request.UserPhotoReq
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	userPhotoReq.UserId, _ = strconv.ParseInt(userID, 10, 64)
	err := ctx.ShouldBind(&userPhotoReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
	}
	switch userPhotoReq.UserPhoto.Header["Content-Type"][0] {
	case "image/png":
	case "image/jpg":
	case "image/jpeg":
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "File tidak sesuai",
		})
		return
	}

	res, err := c.userService.UpdatePhoto(userPhotoReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	path := fmt.Sprintf("src/images/profiles/%s", res.UserPhoto)
	if err := ctx.SaveUploadedFile(userPhotoReq.UserPhoto, path); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Gagal upload foto :" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Berhasil mengubah foto profile",
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
