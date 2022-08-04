package database

import "time"

type Transaction struct {
	TransactionID       int64     `gorm:"primary_key:auto_increment" json:"transaction_id"`
	TransactionUserID   int64     `gorm:"not null" json:"-"`
	TransactionWalletID int64     `gorm:"not null" json:"-"`
	TransactionType     string    `gorm:"type:varchar(100)" json:"transaction_type"`
	TransactionCategory int64     `gorm:"not null" json:"-"`
	TransactionDetail   string    `gorm:"type:text" json:"transaction_detail"`
	TransactionPrice    int64     `json:"transaction_price"`
	TransactionDate     time.Time `json:"transaction_date"`
	User                Users     `gorm:"foreignkey:TransactionUserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"transaction_user_id"`
	Wallet              Wallets   `gorm:"foreignkey:TransactionWalletID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"transaction_wallet_id"`
	Category            Category  `gorm:"foreignkey:TransactionCategory;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"transaction_category_id"`
}
