package services

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type UserService interface {
	CreatedPassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error
	GetUserProfile(userId int64) (database.Users, error)
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

func (s *userService) GetUserProfile(userId int64) (database.Users, error) {
	res, err := s.userRepo.GetUserProfile(userId)
	if err != nil {
		return database.Users{}, err
	}
	return res, nil

}
