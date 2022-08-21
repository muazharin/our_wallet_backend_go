package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type UserService interface {
	CreatedPassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error
	GetUserProfile(userId int64) (response.UserProfileRes, error)
	UpdatePhoto(userPhotoReq request.UserPhotoReq) (database.Users, error)
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreatedPassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error {
	err := s.userRepo.CreatePassword(userCreatePasswordRequest, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserProfile(userId int64) (response.UserProfileRes, error) {
	var userProfileRes response.UserProfileRes
	res, err := s.userRepo.GetUserProfile(userId)
	if err != nil {
		return response.UserProfileRes{}, err
	}
	userProfileRes.UserID = res.UserID
	userProfileRes.UserName = res.UserName
	userProfileRes.UserEmail = res.UserEmail
	userProfileRes.UserPhone = res.UserPhone
	userProfileRes.UserPhoto = fmt.Sprintf("%v/images/profiles/%v", os.Getenv("BASE_URL"), res.UserPhoto)
	userProfileRes.UserGender = res.UserGender
	userProfileRes.UserTglLahir = res.UserTglLahir.Format("2006-01-02 15:04:05")
	userProfileRes.UserAddress = res.UserAddress
	userProfileRes.UserCreatedAt = res.UserCreatedAt.Format("2006-01-02 15:04:05")
	userProfileRes.UserUpdatedAt = res.UserUpdatedAt.Format("2006-01-02 15:04:05")
	return userProfileRes, nil
}

func (s *userService) UpdatePhoto(userPhotoReq request.UserPhotoReq) (database.Users, error) {
	res, err := s.userRepo.GetUserProfile(userPhotoReq.UserId)
	if err != nil {
		return database.Users{}, err
	}
	if res.UserPhoto != "" {
		path := fmt.Sprintf("src/images/profiles/%v", res.UserPhoto)
		err = os.Remove(path)
		if err != nil {
			return database.Users{}, err
		}
	}
	res.UserPhoto = fmt.Sprintf("%v.%v", time.Now().Unix(), strings.Split(userPhotoReq.UserPhoto.Filename, ".")[1])
	res.UserUpdatedAt = time.Now()
	res, err = s.userRepo.UpdatePhoto(res)
	if err != nil {
		return database.Users{}, err
	}

	return res, nil
}
