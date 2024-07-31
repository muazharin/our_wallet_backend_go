package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/services"
	"google.golang.org/api/option"
)

type OWController interface {
	GetOwUser(ctx *gin.Context)
	GetForMember(ctx *gin.Context)
	AddMember(ctx *gin.Context)
	RemoveMember(ctx *gin.Context)
	ConfirmInvitation(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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

func (c *owController) AddMember(ctx *gin.Context) {
	var owAddMemberReq request.OwAddMemberReq
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err := ctx.ShouldBind(&owAddMemberReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	res, err := c.owService.AddMember(owAddMemberReq, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	credential := os.Getenv("CREDENTIAL_FIREBASE")
	project_id := os.Getenv("PROJECT_ID_FIREBASE")
	opt := option.WithCredentialsFile(credential)
	config := &firebase.Config{ProjectID: project_id}
	app, _ := firebase.NewApp(context.Background(), config, opt)
	fcmC, _ := app.Messaging(context.Background())

	var msgs []*messaging.Message
	for i := 0; i < len(res); i++ {
		msgs = append(msgs, &messaging.Message{
			Notification: &messaging.Notification{
				Title: fmt.Sprintf("%v", "Notifikasi"),
				Body:  fmt.Sprintf("%v", "Anda memiliki 1 undangan menjadi anggota wallet"),
			},
			Token: res[i].FirebaseTokenString,
		})
	}
	response, _ := fcmC.SendAll(context.Background(), msgs)
	fmt.Println(response)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User berhasil ditambahkan. Sedang menunggu konfirmasi.",
	})
}

func (c *owController) RemoveMember(ctx *gin.Context) {
	var owAddMemberReq request.OwAddMemberReq
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err := ctx.ShouldBind(&owAddMemberReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	err = c.owService.RemoveMember(owAddMemberReq, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User berhasil dihapus",
	})
}

func (c *owController) ConfirmInvitation(ctx *gin.Context) {
	var confirmInvitation request.OwConfirmInvitation
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Split(authHeader, "Bearer ")[1]
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, _ := strconv.ParseInt(userID, 10, 64)
	err := ctx.Bind(&confirmInvitation)
	if err != nil {
		fmt.Println(confirmInvitation.ConfirmReply)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	err = c.owService.ConfirmInvitation(confirmInvitation, convertedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Undangan berhasil dikonfirmasi",
	})
}

func (c *owController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
