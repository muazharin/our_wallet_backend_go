package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"gorm.io/gorm"
)

type WalletRepo interface {
	GetAllWallet(userId int64, page int64) ([]database.Wallets, error)
	GetWalletById(walletId int64) (database.Wallets, error)
	GetInvitationWallet(userId int64, page int64) ([]database.Wallets, error)
	CreateWallet(wallet database.Wallets, our_wallet database.OurWallet) error
	UpdateWallet(wallet database.Wallets, userId int64, isAdmin bool) error
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
		Where("our_wallets.ow_user_id = ? AND wallets.wallet_is_active = ? AND our_wallets.ow_is_user_active = ?", userId, true, 1).
		Offset((int(page) - 1) * 10).Limit(10).
		Scan(&wallet)
	if err.Error != nil {
		return nil, err.Error
	}
	return wallet, nil
}

func (db *walletConnection) GetWalletById(walletId int64) (database.Wallets, error) {
	var wallet database.Wallets
	res := db.connection.Model(&database.Wallets{}).Where("wallet_id = ?", walletId).First(&wallet)

	if res.Error != nil {
		return database.Wallets{}, res.Error
	}

	return wallet, nil
}

func (db *walletConnection) GetInvitationWallet(userId int64, page int64) ([]database.Wallets, error) {
	var wallet []database.Wallets
	err := db.connection.Model(&database.Wallets{}).
		Joins("left join our_wallets on our_wallets.ow_wallet_id = wallets.wallet_id").
		Where("our_wallets.ow_user_id = ? AND wallets.wallet_is_active = ? AND our_wallets.ow_is_user_active = ?", userId, true, 0).
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

func (db *walletConnection) UpdateWallet(wallet database.Wallets, userId int64, isAdmin bool) error {
	var ow database.OurWallet
	res := db.connection.Where(&database.OurWallet{
		OwWalletID:     wallet.WalletID,
		OwUserID:       userId,
		OwIsAdmin:      isAdmin,
		OwIsUserActive: 1,
	}).First(&ow)

	if res.Error != nil {
		res.Error = fmt.Errorf("anda tidak memiliki hak akses")
		return res.Error
	}

	res = db.connection.Save(&wallet)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal mengupdate wallet")
		return res.Error
	}
	return nil

}
