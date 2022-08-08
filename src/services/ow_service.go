package services

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type OWService interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error)
}

type owService struct {
	owRepo repositories.OWRepo
}

func NewOWService(owRepo repositories.OWRepo) OWService {
	return &owService{
		owRepo: owRepo,
	}
}

func (s *owService) GetOwUser(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error) {
	var getOwUserRes response.GetOwUserRes
	var getOwUserRess []response.GetOwUserRes
	res, err := s.owRepo.GetOwUser(owGetUserReq)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		getOwUserRes.UserID = v.UserID
		getOwUserRes.UserName = v.UserName
		getOwUserRes.UserEmail = v.UserEmail
		getOwUserRes.UserPhone = v.UserPhone
		getOwUserRes.UserPhoto = v.UserPhoto
		getOwUserRes.UserGender = v.UserGender
		getOwUserRes.UserTglLahir = v.UserTglLahir.Format("2006-01-02")
		getOwUserRes.UserAddress = v.UserAddress
		getOwUserRes.UserStatus = v.UserStatus
		getOwUserRess = append(getOwUserRess, getOwUserRes)
	}
	return getOwUserRess, nil
}
