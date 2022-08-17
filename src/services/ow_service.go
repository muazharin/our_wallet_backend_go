package services

import (
	"fmt"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type OWService interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error)
	GetForMember(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error)
	AddMember(owAddMemberReq request.OwAddMemberReq, userId int64) error
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
	res, err := s.owRepo.GetOwUser(owGetUserReq, false)
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

func (s *owService) GetForMember(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error) {
	var getOwUserRes response.GetOwUserRes
	var getOwUserRess []response.GetOwUserRes
	res, err := s.owRepo.GetForMember(owGetUserReq)

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

func (s *owService) AddMember(owAddMemberReq request.OwAddMemberReq, userId int64) error {
	var owWallet database.OurWallet
	count, err := s.owRepo.CheckMember(owAddMemberReq, userId)
	if err != nil || count == 0 {
		err = fmt.Errorf("anda tidak memiliki hak akses")
		return err
	}
	owWallet.OwID = time.Now().Unix()
	owWallet.OwUserID = owAddMemberReq.OwMemberId
	owWallet.OwWalletID = owAddMemberReq.OwWalletId
	owWallet.OwIsUserActive = 0
	owWallet.OwIsAdmin = false
	owWallet.OwDate = time.Now()
	err = s.owRepo.AddMember(owWallet)
	if err != nil {
		return err
	}
	return nil
}
