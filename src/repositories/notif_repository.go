package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"gorm.io/gorm"
)

type NotifRepo interface {
	SendNotif(notifSend database.Notification) error
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
