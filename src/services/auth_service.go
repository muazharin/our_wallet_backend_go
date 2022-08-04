package services

import (
	"fmt"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type AuthService interface {
	CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, error)
	SignUp(authSignUpRequest request.AuthSignUpRequest) (response.AuthSignUpResponse, error)
}

type authService struct {
	authRepo repositories.AuthRepo
}

func NewAuthService(authRepo repositories.AuthRepo) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, error) {
	res, err := s.authRepo.CheckPhone(checkPhoneRequest)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *authService) SignUp(authSignUpRequest request.AuthSignUpRequest) (response.AuthSignUpResponse, error) {
	authSignUpResponse := response.AuthSignUpResponse{}
	user := database.Users{}
	count, err := s.authRepo.CheckAccount(authSignUpRequest.UserName, authSignUpRequest.UserEmail, authSignUpRequest.UserPhone)
	if err != nil {
		return authSignUpResponse, err
	}
	if count > 0 {
		err = fmt.Errorf("Username, Nomor Hp atau Email sudah terdaftar")
		return authSignUpResponse, err
	}
	layout := "2006-01-02"
	d, err := time.Parse(layout, authSignUpRequest.UserTglLahir)
	if err != nil {
		return authSignUpResponse, err
	}
	user.UserID = time.Now().Unix()
	user.UserName = authSignUpRequest.UserName
	user.UserEmail = authSignUpRequest.UserEmail
	user.UserPhone = authSignUpRequest.UserPhone
	user.UserGender = authSignUpRequest.UserGender
	user.UserTglLahir = d
	user.UserAddress = authSignUpRequest.UserAddress
	user.UserCreatedAt = time.Now()
	user.UserUpdatedAt = time.Now()
	err = s.authRepo.SignUp(user)
	if err != nil {
		return authSignUpResponse, err
	}
	authSignUpResponse.UserID = time.Now().Unix()
	authSignUpResponse.UserName = authSignUpRequest.UserName
	authSignUpResponse.UserEmail = authSignUpRequest.UserEmail
	authSignUpResponse.UserPhone = authSignUpRequest.UserPhone
	authSignUpResponse.UserGender = authSignUpRequest.UserGender
	authSignUpResponse.UserTglLahir = authSignUpRequest.UserTglLahir
	authSignUpResponse.UserAddress = authSignUpRequest.UserAddress
	authSignUpResponse.UserCreatedAt = fmt.Sprintf("%v", user.UserCreatedAt.Format("2006-01-02 15:04:05"))
	authSignUpResponse.UserUpdatedAt = fmt.Sprintf("%v", user.UserUpdatedAt.Format("2006-01-02 15:04:05"))

	return authSignUpResponse, nil
}
