package services

import (
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
	err := s.userRepo.CreatePassword(userCreatePasswordRequest, userId)
	if err != nil {
		return err
	}
	return nil
}
