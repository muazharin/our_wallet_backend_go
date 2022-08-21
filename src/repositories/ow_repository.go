package repositories

import (
	"fmt"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type OWRepo interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]database.Users, error)
	GetForMember(owGetUserReq request.OwGetUserReq) ([]database.Users, error)
	CheckMember(owAddMemberReq request.OwAddMemberReq, userId int64) (int64, error)
	AddMember(owWallet database.OurWallet) error
	RemoveMember(owAddMemberReq request.OwAddMemberReq) (string, error)
	ConfirmInvitation(confirmInvitation request.OwConfirmInvitation, userId int64) (database.OurWallet, error)
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
		Where("our_wallets.ow_is_user_active = ? AND our_wallets.ow_wallet_id = ?", 1, owGetUserReq.WalletId).
		Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
		Scan(&user)

	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil
}

func (db *owConnection) GetForMember(owGetUserReq request.OwGetUserReq) ([]database.Users, error) {
	var listId []int64
	var users []database.Users
	res, e := db.GetOwUser(owGetUserReq)
	if e != nil {
		return nil, e
	}
	for _, v := range res {
		listId = append(listId, v.UserID)
	}
	fmt.Println(listId)
	fmt.Println(res)
	keyword := fmt.Sprintf("%v", owGetUserReq.Keyword)
	var err *gorm.DB
	if owGetUserReq.Keyword != "" {
		err = db.connection.Model(&database.Users{}).
			Or("user_email LIKE ?", "%"+keyword+"%").
			Or("user_phone LIKE ?", "%"+keyword+"%").
			Where("user_id NOT IN ? AND user_name LIKE ?", listId, "%"+keyword+"%").
			Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
			Scan(&users)
	} else {
		err = db.connection.Model(&database.Users{}).
			Where("user_id NOT IN ?", listId).
			Offset((int(owGetUserReq.Page) - 1) * 10).Limit(10).
			Scan(&users)
		fmt.Println(users)
	}
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}

func (db *owConnection) CheckMember(owAddMemberReq request.OwAddMemberReq, userId int64) (int64, error) {
	var count int64
	err := db.connection.Model(&database.OurWallet{}).
		Where("ow_wallet_id=? AND ow_is_admin=? AND ow_user_id=?", owAddMemberReq.OwWalletId, true, userId).
		Count(&count)
	if err.Error != nil {
		return count, err.Error
	}
	return count, nil
}

func (db *owConnection) AddMember(owWallet database.OurWallet) error {
	var ow database.OurWallet
	var count int64
	db.connection.Model(&database.OurWallet{}).
		Where(&database.OurWallet{
			OwWalletID: owWallet.OwWalletID,
			OwUserID:   owWallet.OwUserID,
		}).First(&ow).Count(&count)
	if count > 0 {
		switch ow.OwIsUserActive {
		case 0:
			err := fmt.Errorf("user telah diundang, sedang menunggu konfirmasi")
			return err
		case 2:
			ow.OwIsUserActive = 1
			db.connection.Save(&ow)
			return nil
		}
	}
	res := db.connection.Save(&owWallet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *owConnection) RemoveMember(owAddMemberReq request.OwAddMemberReq) (string, error) {
	var owWallet database.OurWallet
	var wallet database.Wallets
	res := db.connection.Where(&database.OurWallet{
		OwWalletID: owAddMemberReq.OwWalletId,
		OwUserID:   owAddMemberReq.OwMemberId,
	}).First(&owWallet)
	if res.Error != nil {
		return "", res.Error
	}
	res = db.connection.Where(&database.Wallets{
		WalletID: owWallet.OwWalletID,
	}).First(&wallet)
	if res.Error != nil {
		return "", res.Error
	}
	owWallet.OwIsUserActive = 2
	res = db.connection.Save(&owWallet)
	if res.Error != nil {
		return "", res.Error
	}
	return wallet.WalletName, nil
}

func (db *owConnection) ConfirmInvitation(confirmInvitation request.OwConfirmInvitation, userId int64) (database.OurWallet, error) {
	var owWallet database.OurWallet
	err := db.connection.
		Where("ow_wallet_id = ? AND ow_user_id = ?", confirmInvitation.ConfirmWalletId, userId).
		First(&owWallet)
	if err.Error != nil {
		return owWallet, err.Error
	}
	if confirmInvitation.ConfirmReply {
		owWallet.OwIsUserActive = 1
		owWallet.OwDate = time.Now()
		err = db.connection.Save(&owWallet)
	} else {
		err = db.connection.Delete(&owWallet)
	}

	if err.Error != nil {
		err.Error = fmt.Errorf("gagal mengkonfirmasi undangan")
		return owWallet, err.Error
	}
	return owWallet, nil

}
