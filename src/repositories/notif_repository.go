package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type NotifRepo interface {
	SendNotif(notifSend database.Notification) error
	GetAllNotif(userId int64, page int64) ([]database.Notification, error)
	IsReadNotif(isReadNotifReq request.IsReadNotifReq) error
}

type notifRepo struct {
	connection *gorm.DB
}

func NewNotifRepo(connection *gorm.DB) NotifRepo {
	return &notifRepo{
		connection: connection,
	}
}

func (db *notifRepo) SendNotif(notifSend database.Notification) error {
	err := db.connection.Save(&notifSend)
	if err.Error != nil {
		err.Error = fmt.Errorf("gagal mengirim notifikasi")
		return err.Error
	}
	return nil
}

func (db *notifRepo) GetAllNotif(userId int64, page int64) ([]database.Notification, error) {
	var notif database.Notification
	var notifs []database.Notification
	err := db.connection.Where(&database.Notification{
		NotificationReceiverID: userId,
	}).Offset((int(page) - 1) * 10).Limit(10).
		Find(&notif).
		Scan(&notifs)
	if err.Error != nil {
		return nil, err.Error
	}

	return notifs, nil
}

func (db *notifRepo) IsReadNotif(isReadNotifReq request.IsReadNotifReq) error {
	var notif database.Notification
	err := db.connection.Where(&database.Notification{
		NotificationID: isReadNotifReq.NotifId,
	}).First(&notif)
	if err.Error != nil {
		err.Error = fmt.Errorf("notifikasi tidak ditemukan")
		return err.Error
	}
	notif.NotificationIsRead = true
	err = db.connection.Save(&notif)
	if err.Error != nil {
		err.Error = fmt.Errorf("gagal membaca notifikasi")
		return err.Error
	}
	return nil
}
