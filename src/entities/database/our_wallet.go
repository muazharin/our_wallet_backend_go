package database

import "time"

type OurWallet struct {
	OwID           int64     `gorm:"primary_key:auto_increment" json:"ow_id"`
	OwWalletID     int64     `gorm:"not null" json:"ow_wallet_id"`
	OwUserID       int64     `gorm:"not null" json:"ow_user_id"`
	OwIsUserActive int       `gorm:"not null" json:"ow_is_user_active"`
	OwIsAdmin      bool      `gorm:"not null" json:"ow_is_admin"`
	OwDate         time.Time `gorm:"not null" json:"ow_date"`
	OwInviterID    int64     `gorm:"not null" json:"ow_inviter_id"`
}
