package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type OWRepo interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]database.Users, error)
	GetForMember(owGetUserReq request.OwGetUserReq) ([]database.Users, error)
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

func (db *owConnection) GetForMember(owGetUserReq request.OwGetUserReq) ([]database.Users, error) {
	var listId []int64
	var user database.Users
	var users []database.Users
	res, e := db.GetOwUser(owGetUserReq)
	if e != nil {
		return nil, e
	}
	for _, v := range res {
		listId = append(listId, v.UserID)
	}
	fmt.Println(res)
	keyword := fmt.Sprintf("%v", owGetUserReq.Keyword)
	var err *gorm.DB
	if owGetUserReq.Keyword != "" {
		err = db.connection.
			Or("user_email LIKE ?", "%"+keyword+"%").
			Or("user_phone LIKE ?", "%"+keyword+"%").
			Where("user_id NOT IN (?) AND user_name LIKE ?", listId, "%"+keyword+"%").
			Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
			Find(&user).
			Scan(&users)
	} else {
		err = db.connection.Not(listId).
			Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
			Find(&user).Scan(&users)
	}
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}
