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

type TransService interface {
	CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error)
	GetAllTransByWalletId(transWalletId request.TransByWalletIdReq) ([]response.TransByWalletIdRes, error)
	GetAllTransByUserId(transUserId request.TransByUserIdReq) ([]response.TransByUserIdRes, error)
	GetTransById(transId request.TransByIdReq) (response.TransByIdRes, error)
}

type transService struct {
	transRepo    repositories.TransRepo
	categoryRepo repositories.CategoryRepo
	userRepo     repositories.UserRepo
	walletRepo   repositories.WalletRepo
}

func NewTransService(transRepo repositories.TransRepo, categoryRepo repositories.CategoryRepo, userRepo repositories.UserRepo, walletRepo repositories.WalletRepo) TransService {
	return &transService{
		transRepo:    transRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
		walletRepo:   walletRepo,
	}
}

func (s *transService) CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error) {
	var transaction database.Transaction
	var trFile database.TransactionFile
	var trFiles []database.TransactionFile
	transaction.TransactionID = time.Now().Unix()
	transaction.TransactionUserID = userId
	transaction.TransactionWalletID = transCreateReq.TransWalletId
	transaction.TransactionType = transCreateReq.TransType
	transaction.TransactionCategory = transCreateReq.TransCategoryId
	transaction.TransactionDetail = transCreateReq.TransDetail
	transaction.TransactionPrice = transCreateReq.TransPrice
	transaction.TransactionDate = time.Now()
	err := s.transRepo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	for i, v := range transCreateReq.TransFile {
		trFile.TfID = time.Now().Unix() + int64(i)
		trFile.TfTransactionID = transaction.TransactionID
		trFile.TfFile = fmt.Sprintf("%v%v.%v", time.Now().Unix(), i, strings.Split(v.Filename, ".")[1])
		trFiles = append(trFiles, trFile)
		err = s.transRepo.SaveFileTrans(trFile)
		if err != nil {
			err = fmt.Errorf("gagal menyimpan file transaksi")
			return nil, err
		}
	}
	wallet, err := s.walletRepo.GetWalletById(transCreateReq.TransWalletId)
	if err != nil {
		return nil, err
	}
	wallet.WalletMoney = (wallet.WalletMoney - transCreateReq.TransPrice)
	err = s.walletRepo.UpdateWallet(wallet, userId, false)
	if err != nil {
		return nil, err
	}

	return trFiles, nil
}

func (s *transService) GetAllTransByWalletId(transWalletId request.TransByWalletIdReq) ([]response.TransByWalletIdRes, error) {
	var transByWalletIdRes response.TransByWalletIdRes
	var transByWalletIdRess []response.TransByWalletIdRes
	res, err := s.transRepo.GetAllTransByWalletId(transWalletId)
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		users, _ := s.userRepo.GetUserProfile(v.TransactionUserID)
		category, _ := s.categoryRepo.GetCategoryById(v.TransactionCategory)
		isSeen, _ := s.transRepo.CheckIsSeen(v.TransactionID, v.TransactionUserID)
		transByWalletIdRes.TransID = v.TransactionID
		transByWalletIdRes.TransType = v.TransactionType
		transByWalletIdRes.TransCategory = category.CategoryTitle
		transByWalletIdRes.TransDetail = v.TransactionDetail
		transByWalletIdRes.TransPrice = v.TransactionPrice
		transByWalletIdRes.TransDate = v.TransactionDate.Format("2006-01-02 15:04:05")
		transByWalletIdRes.TransIsSeen = isSeen
		transByWalletIdRes.TransUser.TransUserID = users.UserID
		transByWalletIdRes.TransUser.TransUserName = users.UserName
		transByWalletIdRes.TransUser.TransUserEmail = users.UserEmail
		transByWalletIdRes.TransUser.TransUserPhone = users.UserPhone
		transByWalletIdRes.TransUser.TransUserPhoto = users.UserPhoto
		if users.UserPhoto != "" {
			transByWalletIdRes.TransUser.TransUserPhoto = fmt.Sprintf("%v/images/profiles/%v", os.Getenv("BASE_URL"), users.UserPhoto)

		}
		transByWalletIdRess = append(transByWalletIdRess, transByWalletIdRes)

	}
	return transByWalletIdRess, nil
}

func (s *transService) GetAllTransByUserId(transUserId request.TransByUserIdReq) ([]response.TransByUserIdRes, error) {
	var transByUserIdRes response.TransByUserIdRes
	var transByUserIdRess []response.TransByUserIdRes
	res, err := s.transRepo.GetAllTransByUserId(transUserId)
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		wallet, _ := s.walletRepo.GetWalletById(v.TransactionWalletID)
		category, _ := s.categoryRepo.GetCategoryById(v.TransactionCategory)
		isSeen, _ := s.transRepo.CheckIsSeen(v.TransactionID, v.TransactionUserID)
		transByUserIdRes.TransID = v.TransactionID
		transByUserIdRes.TransType = v.TransactionType
		transByUserIdRes.TransCategory = category.CategoryTitle
		transByUserIdRes.TransDetail = v.TransactionDetail
		transByUserIdRes.TransPrice = v.TransactionPrice
		transByUserIdRes.TransDate = v.TransactionDate.Format("2006-01-02 15:04:05")
		transByUserIdRes.TransIsSeen = isSeen
		transByUserIdRes.TransWallet.TransWalletID = wallet.WalletID
		transByUserIdRes.TransWallet.TransWalletName = wallet.WalletName
		transByUserIdRes.TransWallet.TransWalletColor = wallet.WalletColor
		transByUserIdRess = append(transByUserIdRess, transByUserIdRes)
	}
	return transByUserIdRess, nil
}

func (s *transService) GetTransById(transId request.TransByIdReq) (response.TransByIdRes, error) {
	var transByIdRes response.TransByIdRes
	var transIsSeen database.TransactionIsSeen
	res, err := s.transRepo.GetTransById(transId)
	if err != nil {
		return response.TransByIdRes{}, err
	}
	isSeen, _ := s.transRepo.CheckIsSeen(res.TransactionID, res.TransactionUserID)
	if !isSeen {
		transIsSeen.TransactionIsSeenID = time.Now().Unix()
		transIsSeen.TransactionID = res.TransactionID
		transIsSeen.TransactionUserID = res.TransactionUserID
		s.transRepo.SetIsSeen(transIsSeen)
	}
	users, _ := s.userRepo.GetUserProfile(res.TransactionUserID)
	wallet, _ := s.walletRepo.GetWalletById(res.TransactionWalletID)
	category, _ := s.categoryRepo.GetCategoryById(res.TransactionCategory)
	trFile := response.TransFile{}
	trFiles, _ := s.transRepo.GetFileByTransId(res.TransactionID, 0)
	transByIdRes.TransID = res.TransactionID
	transByIdRes.TransType = res.TransactionType
	transByIdRes.TransCategory = category.CategoryTitle
	transByIdRes.TransDetail = res.TransactionDetail
	transByIdRes.TransPrice = res.TransactionPrice
	transByIdRes.TransDate = res.TransactionDate.Format("2006-01-02 15:04:05")
	transByIdRes.TransIsSeen = true
	transByIdRes.TransUser.TransUserID = users.UserID
	transByIdRes.TransUser.TransUserName = users.UserName
	transByIdRes.TransUser.TransUserEmail = users.UserEmail
	transByIdRes.TransUser.TransUserPhone = users.UserPhone
	transByIdRes.TransUser.TransUserPhoto = users.UserPhoto
	if users.UserPhoto != "" {
		transByIdRes.TransUser.TransUserPhoto = fmt.Sprintf("%v/images/profiles/%v", os.Getenv("BASE_URL"), users.UserPhoto)
	}
	transByIdRes.TransWallet.TransWalletID = wallet.WalletID
	transByIdRes.TransWallet.TransWalletName = wallet.WalletName
	transByIdRes.TransWallet.TransWalletColor = wallet.WalletColor
	transByIdRes.TransFile = []response.TransFile{}
	if len(trFiles) != 0 {
		for _, v := range trFiles {
			trFile.TransFileID = v.TfID
			trFile.TransFileImage = fmt.Sprintf("%v/images/trFiles/%v", os.Getenv("BASE_URL"), v.TfFile)
			transByIdRes.TransFile = append(transByIdRes.TransFile, trFile)
		}
	}

	return transByIdRes, nil
}
