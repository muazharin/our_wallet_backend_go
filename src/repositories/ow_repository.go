package repositories

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type OWRepo interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]database.Users, error)
}

type owConnection struct {
	connection *gorm.DB
}

func NewOWRepo(connection *gorm.DB) OWRepo {
	return &owConnection{
		connection: connection,
	}
}

func (db *owConnection) GetOwUser(owGetUserReq request.OwGetUserReq) ([]database.Users, error) {
	var user []database.Users
	err := db.connection.Model(&database.Users{}).
		Joins("left join our_wallets ON our_wallets.ow_user_id = users.user_id").
		Where("our_wallets.ow_is_user_active = ? AND our_wallets.ow_wallet_id = ?", true, owGetUserReq.WalletId).
		Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
		Scan(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil

}
