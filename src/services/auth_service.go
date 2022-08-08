package services

import (
	"fmt"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, response.AuthSignUpResponse, error)
	SignUp(authSignUpRequest request.AuthSignUpRequest) (response.AuthSignUpResponse, error)
	SignIn(authSignInRequest request.AuthSignInRequest) (response.AuthSignUpResponse, error)
}

type authService struct {
	authRepo repositories.AuthRepo
}

func NewAuthService(authRepo repositories.AuthRepo) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, response.AuthSignUpResponse, error) {
	authSignUpResponse := response.AuthSignUpResponse{}
	count, res, err := s.authRepo.CheckAccount(checkPhoneRequest.Phone, checkPhoneRequest.Phone, checkPhoneRequest.Phone)
	if err != nil {
		return count, response.AuthSignUpResponse{}, err
	}
	authSignUpResponse.UserID = time.Now().Unix()
	authSignUpResponse.UserName = res.UserName
	authSignUpResponse.UserEmail = res.UserEmail
	authSignUpResponse.UserPhone = res.UserPhone
	authSignUpResponse.UserGender = res.UserGender
	authSignUpResponse.UserTglLahir = res.UserTglLahir.Format("2006-01-02")
	authSignUpResponse.UserAddress = res.UserAddress
	authSignUpResponse.UserStatus = res.UserStatus
	authSignUpResponse.UserCreatedAt = fmt.Sprintf("%v", res.UserCreatedAt.Format("2006-01-02 15:04:05"))
	authSignUpResponse.UserUpdatedAt = fmt.Sprintf("%v", res.UserUpdatedAt.Format("2006-01-02 15:04:05"))
	return count, authSignUpResponse, nil
}

func (s *authService) SignUp(authSignUpRequest request.AuthSignUpRequest) (response.AuthSignUpResponse, error) {
	authSignUpResponse := response.AuthSignUpResponse{}
	user := database.Users{}
	count, _, err := s.authRepo.CheckAccount(authSignUpRequest.UserName, authSignUpRequest.UserEmail, authSignUpRequest.UserPhone)
	if err != nil {
		return authSignUpResponse, err
	}
	if count > 0 {
		err = fmt.Errorf("username, nomor hp atau email sudah terdaftar")
		return authSignUpResponse, err
	}
	layout := "2006-01-02"
	tglLahir, err := time.Parse(layout, authSignUpRequest.UserTglLahir)
	if err != nil {
		return authSignUpResponse, err
	}
	user.UserID = time.Now().Unix()
	user.UserName = authSignUpRequest.UserName
	user.UserEmail = authSignUpRequest.UserEmail
	user.UserPhone = authSignUpRequest.UserPhone
	user.UserGender = authSignUpRequest.UserGender
	user.UserTglLahir = tglLahir
	user.UserAddress = authSignUpRequest.UserAddress
	user.UserStatus = "incomplete"
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
	authSignUpResponse.UserStatus = "incomplete"
	authSignUpResponse.UserCreatedAt = fmt.Sprintf("%v", user.UserCreatedAt.Format("2006-01-02 15:04:05"))
	authSignUpResponse.UserUpdatedAt = fmt.Sprintf("%v", user.UserUpdatedAt.Format("2006-01-02 15:04:05"))

	return authSignUpResponse, nil
}

func (s *authService) SignIn(authSignInRequest request.AuthSignInRequest) (response.AuthSignUpResponse, error) {
	authSignUpResponse := response.AuthSignUpResponse{}
	res, err := s.authRepo.SignIn(authSignInRequest)
	if err != nil {
		fmt.Println("1")
		return authSignUpResponse, err
	}
	compared, err := comparePassword(res.UserPassword, []byte(authSignInRequest.UserPassword))
	if err != nil {
		err = fmt.Errorf("password salah")
		return authSignUpResponse, err
	}

	if (res.UserName == authSignInRequest.UserName || res.UserEmail == authSignInRequest.UserName || res.UserPhone == authSignInRequest.UserName) && compared {
		authSignUpResponse.UserID = res.UserID
		authSignUpResponse.UserName = res.UserName
		authSignUpResponse.UserPassword = res.UserPassword
		authSignUpResponse.UserEmail = res.UserEmail
		authSignUpResponse.UserPhone = res.UserPhone
		authSignUpResponse.UserPhoto = res.UserPhoto
		authSignUpResponse.UserGender = res.UserGender
		authSignUpResponse.UserTglLahir = res.UserTglLahir.Format("2006-01-02")
		authSignUpResponse.UserAddress = res.UserAddress
		authSignUpResponse.UserStatus = res.UserStatus
		authSignUpResponse.UserCreatedAt = res.UserCreatedAt.Format("2006-01-02 15:04:05")
		authSignUpResponse.UserUpdatedAt = res.UserUpdatedAt.Format("2006-01-02 15:04:05")
		return authSignUpResponse, nil
	}

	err = fmt.Errorf("user tidak ditemukan")
	return authSignUpResponse, err
}

func comparePassword(hashedPwd string, plainPassword []byte) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
