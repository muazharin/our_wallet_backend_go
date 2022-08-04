package database

import "time"

type OurWallet struct {
	OwID           int64     `gorm:"primary_key:auto_increment" json:"ow_id"`
	OwWalletID     int64     `gorm:"not null" json:"-"`
	OwUserID       int64     `gorm:"not null" json:"-"`
	OwIsUserActive bool      `json:"ow_is_user_active"`
	OwIsAdmin      bool      `json:"ow_is_admin"`
	OwDate         time.Time `json:"ow_date"`
	Wallet         Wallets   `gorm:"foreignkey:OwWalletID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"ow_wallet_id"`
	User           Users     `gorm:"foreignkey:OwUserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"ow_user_id"`
}
