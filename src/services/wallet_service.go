package services

import (
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type WalletService interface {
	GetAllWallet(userId int64, page int64) ([]database.Wallets, error)
	CreateWallet(createwallet request.WalletCreateReq, userId int64) error
}

type walletService struct {
	walletRepo repositories.WalletRepo
}

func NewWalletService(walletRepo repositories.WalletRepo) WalletService {
	return &walletService{
		walletRepo: walletRepo,
	}
}

func (s *walletService) GetAllWallet(userId int64, page int64) ([]database.Wallets, error) {
	res, err := s.walletRepo.GetAllWallet(userId, page)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *walletService) CreateWallet(createwallet request.WalletCreateReq, userId int64) error {
	var wallet database.Wallets
	var our_wallet database.OurWallet
	wallet.WalletID = time.Now().Unix()
	wallet.WalletName = createwallet.Name
	wallet.WalletMoney = createwallet.Money
	wallet.WalletColor = createwallet.Color
	wallet.WalletCreatedAt = time.Now()
	wallet.WalletUpdatedAt = time.Now()
	wallet.WalletIsActive = true
	our_wallet.OwID = time.Now().Unix()
	our_wallet.OwWalletID = wallet.WalletID
	our_wallet.OwUserID = userId
	our_wallet.OwIsUserActive = 1
	our_wallet.OwIsAdmin = true
	our_wallet.OwDate = time.Now()
	err := s.walletRepo.CreateWallet(wallet, our_wallet)
	if err != nil {
		return err
	}
	return nil
}
