package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
)

type JWTService interface {
	GenerateToken(v response.AuthSignUpResponse) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClaim struct {
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	UserEmail     string `json:"user_email"`
	UserPhone     string `json:"user_phone"`
	UserPhoto     string `json:"user_photo"`
	UserGender    string `json:"user_gender"`
	UserTglLahir  string `json:"user_tgl_lahir"`
	UserAddress   string `json:"user_address"`
	UserCreatedAt string `json:"user_created_at"`
	UserUpdatedAt string `json:"user_updated_at"`
	jwt.StandardClaims
}
type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("OW_JWT_SECRET")
	return secretKey
}

func (j *jwtService) GenerateToken(v response.AuthSignUpResponse) string {
	claims := &jwtCustomClaim{
		strconv.Itoa(int(v.UserID)),
		v.UserName,
		v.UserEmail,
		v.UserPhone,
		v.UserPhoto,
		v.UserGender,
		v.UserTglLahir,
		v.UserAddress,
		v.UserCreatedAt,
		v.UserUpdatedAt,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
