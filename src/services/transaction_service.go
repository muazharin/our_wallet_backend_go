package services

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type TransService interface {
	CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error)
}

type transService struct {
	transRepo repositories.TransRepo
}

func NewTransService(transRepo repositories.TransRepo) TransService {
	return &transService{
		transRepo: transRepo,
	}
}

func (s *transService) CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error) {
	res, err := s.transRepo.CreateTransaction(transCreateReq, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
