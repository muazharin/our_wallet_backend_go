package database

import "time"

type Wallets struct {
	WalletID        int64        `gorm:"primary_key:auto_increment" json:"wallet_id"`
	WalletName      string       `gorm:"type:varchar(100)" json:"wallet_name"`
	WalletMoney     int64        `gorm:"type:varchar(100)" json:"wallet_money"`
	WalletColor     string       `gorm:"type:varchar(100)" json:"wallet_color"`
	WalletCreatedAt time.Time    `json:"wallet_created_at"`
	WalletUpdatedAt time.Time    `json:"wallet_updated_at"`
	WalletIsActive  bool         `json:"wallet_is_active"`
	WalletOwWallet  []*OurWallet `gorm:"-" json:"wallet_ow_wallet,omitempty"`
}
