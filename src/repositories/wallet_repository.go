package repositories

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"gorm.io/gorm"
)

type WalletRepo interface {
	GetAllWallet(userId int64, page int64) ([]database.Wallets, error)
	CreateWallet(wallet database.Wallets, our_wallet database.OurWallet) error
}

type walletConnection struct {
	connection *gorm.DB
}

func NewWalletRepo(connection *gorm.DB) WalletRepo {
	return &walletConnection{
		connection: connection,
	}
}

func (db *walletConnection) GetAllWallet(userId int64, page int64) ([]database.Wallets, error) {
	var wallet []database.Wallets
	err := db.connection.Model(&database.Wallets{}).
		Joins("left join our_wallets on our_wallets.ow_wallet_id = wallets.wallet_id").
		Where("our_wallets.ow_user_id = ? AND wallets.wallet_is_active = ?", userId, true).
		Offset((int(page) - 1) * 10).Limit(10).
		Scan(&wallet)
	if err.Error != nil {
		return nil, err.Error
	}
	return wallet, nil
}

func (db *walletConnection) CreateWallet(wallet database.Wallets, our_wallet database.OurWallet) error {
	err := db.connection.Save(&wallet)
	if err.Error != nil {
		return err.Error
	}
	err = db.connection.Save(&our_wallet)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
