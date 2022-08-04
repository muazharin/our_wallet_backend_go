package services

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type UserService interface {
	CreatedPassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error
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
	count, err := s.userRepo.CheckUserByID(userId)
	if err != nil {
		return err
	}
	fmt.Println(count)
	if count <= 0 {
		err = fmt.Errorf("User tidak ditemukan")
		return err
	}
	err = s.userRepo.CreatePassword(userCreatePasswordRequest, userId)
	if err != nil {
		err = fmt.Errorf("Password Gagal dibuat")
		return err
	}
	return nil
}
