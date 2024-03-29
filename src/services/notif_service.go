package services

import (
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type NotifService interface {
	GetAllNotif(userId int64, page int64) ([]database.Notification, error)
	IsReadNotif(isReadNotifReq request.IsReadNotifReq) error
}

type notifService struct {
	notifRepo repositories.NotifRepo
}

func NewNotifService(notifRepo repositories.NotifRepo) NotifService {
	return &notifService{
		notifRepo: notifRepo,
	}
}

func (s *notifService) GetAllNotif(userId int64, page int64) ([]database.Notification, error) {
	res, err := s.notifRepo.GetAllNotif(userId, page)
	if err != nil {
		return []database.Notification{}, err
	}
	return res, nil
}

func (s *notifService) IsReadNotif(isReadNotifReq request.IsReadNotifReq) error {
	err := s.notifRepo.IsReadNotif(isReadNotifReq)
	if err != nil {
		return err
	}
	return nil
}
