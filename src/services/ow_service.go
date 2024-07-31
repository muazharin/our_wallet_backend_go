package services

import (
	"fmt"
	"os"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/entities/response"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type OWService interface {
	GetOwUser(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error)
	GetForMember(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error)
	AddMember(owAddMemberReq request.OwAddMemberReq, userId int64) ([]database.FirebaseToken, error)
	RemoveMember(owAddMemberReq request.OwAddMemberReq, userId int64) error
	ConfirmInvitation(confirmInvitation request.OwConfirmInvitation, userId int64) error
}

type owService struct {
	owRepo    repositories.OWRepo
	notifRepo repositories.NotifRepo
}

func NewOWService(owRepo repositories.OWRepo, notifRepo repositories.NotifRepo) OWService {
	return &owService{
		owRepo:    owRepo,
		notifRepo: notifRepo,
	}
}

func (s *owService) GetOwUser(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error) {
	var getOwUserRes response.GetOwUserRes
	var getOwUserRess []response.GetOwUserRes
	res, err := s.owRepo.GetOwUser(owGetUserReq)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		getOwUserRes.UserID = v.UserID
		getOwUserRes.UserName = v.UserName
		getOwUserRes.UserEmail = v.UserEmail
		getOwUserRes.UserPhone = v.UserPhone
		getOwUserRes.UserPhoto = v.UserPhoto
		if v.UserPhoto != "" {
			getOwUserRes.UserPhoto = fmt.Sprintf("%v/images/profiles/%v", os.Getenv("BASE_URL"), v.UserPhoto)
		}
		getOwUserRes.UserGender = v.UserGender
		getOwUserRes.UserTglLahir = v.UserTglLahir.Format("2006-01-02")
		getOwUserRes.UserAddress = v.UserAddress
		getOwUserRes.UserStatus = v.UserStatus
		getOwUserRess = append(getOwUserRess, getOwUserRes)
	}
	return getOwUserRess, nil
}

func (s *owService) GetForMember(owGetUserReq request.OwGetUserReq) ([]response.GetOwUserRes, error) {
	var getOwUserRes response.GetOwUserRes
	var getOwUserRess []response.GetOwUserRes
	res, err := s.owRepo.GetForMember(owGetUserReq)

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		getOwUserRes.UserID = v.UserID
		getOwUserRes.UserName = v.UserName
		getOwUserRes.UserEmail = v.UserEmail
		getOwUserRes.UserPhone = v.UserPhone
		getOwUserRes.UserPhoto = v.UserPhoto
		if v.UserPhoto != "" {
			getOwUserRes.UserPhoto = fmt.Sprintf("%v/images/profiles/%v", os.Getenv("BASE_URL"), v.UserPhoto)
		}
		getOwUserRes.UserGender = v.UserGender
		getOwUserRes.UserTglLahir = v.UserTglLahir.Format("2006-01-02")
		getOwUserRes.UserAddress = v.UserAddress
		getOwUserRes.UserStatus = v.UserStatus
		getOwUserRess = append(getOwUserRess, getOwUserRes)
	}
	return getOwUserRess, nil
}

func (s *owService) AddMember(owAddMemberReq request.OwAddMemberReq, userId int64) ([]database.FirebaseToken, error) {
	var owWallet database.OurWallet
	var notif database.Notification
	var firebaseTokens []database.FirebaseToken

	// mengecek keanggotaan
	count, err := s.owRepo.CheckMember(owAddMemberReq, userId)
	if err != nil || count == 0 {
		err = fmt.Errorf("anda tidak memiliki hak akses")
		return firebaseTokens, err
	}

	// mengajukan undangan
	owWallet.OwID = time.Now().Unix()
	owWallet.OwUserID = owAddMemberReq.OwMemberId
	owWallet.OwWalletID = owAddMemberReq.OwWalletId
	owWallet.OwIsUserActive = 0
	owWallet.OwIsAdmin = false
	owWallet.OwInviterID = userId
	err = s.owRepo.AddMember(owWallet)
	if err != nil {
		return firebaseTokens, err
	}

	// menarik data firebase token
	res, _ := s.owRepo.GetFirebaseTokens(owAddMemberReq.OwMemberId)
	firebaseTokens = res

	// mengirim notifikasi kepada calon member
	notif.NotificationID = time.Now().Unix()
	notif.NotificationPusherID = userId
	notif.NotificationReceiverID = owAddMemberReq.OwMemberId
	notif.NotificationMessage = "Anda memiliki 1 undangan menjadi anggota wallet"
	notif.NotificationRoute = "/invitation"
	notif.NotificationArgument = fmt.Sprintf("{\"wallet_id\":%v}", owWallet.OwWalletID)
	notif.NotificationIsRead = false
	s.notifRepo.SendNotif(notif)
	return firebaseTokens, err
}

func (s *owService) RemoveMember(owAddMemberReq request.OwAddMemberReq, userId int64) error {
	var notif database.Notification
	// mengecek admin
	count, err := s.owRepo.CheckMember(owAddMemberReq, userId)
	if err != nil || count == 0 {
		err = fmt.Errorf("anda tidak memiliki hak akses")
		return err
	}
	// mengeluarkan user
	res, err := s.owRepo.RemoveMember(owAddMemberReq)
	if err != nil {
		return err
	}

	// mengirim notifikasi kepada calon member
	notif.NotificationID = time.Now().Unix()
	notif.NotificationPusherID = userId
	notif.NotificationReceiverID = owAddMemberReq.OwMemberId
	notif.NotificationMessage = fmt.Sprintf("Anda telah dikeluarkan dari wallet %v", res)
	notif.NotificationRoute = "/"
	notif.NotificationIsRead = false
	s.notifRepo.SendNotif(notif)
	return nil
}

func (s *owService) ConfirmInvitation(confirmInvitation request.OwConfirmInvitation, userId int64) error {
	var notif database.Notification
	// mengirim konfirmasi
	ow, err := s.owRepo.ConfirmInvitation(confirmInvitation, userId)
	if err != nil {
		return err
	}

	// mengirim notifikasi
	notif.NotificationID = time.Now().Unix()
	notif.NotificationPusherID = userId
	notif.NotificationReceiverID = ow.OwInviterID
	if confirmInvitation.ConfirmReply {
		notif.NotificationMessage = "Undangan anda diterima"
	} else {
		notif.NotificationMessage = "Undangan anda ditolak"
	}
	notif.NotificationRoute = "/"
	notif.NotificationIsRead = false
	s.notifRepo.SendNotif(notif)
	return nil
}
