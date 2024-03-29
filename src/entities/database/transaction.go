package database

import "time"

type Transaction struct {
	TransactionID       int64     `gorm:"primary_key:auto_increment" json:"transaction_id"`
	TransactionUserID   int64     `gorm:"not null" json:"transaction_user_id"`
	TransactionWalletID int64     `gorm:"not null" json:"transaction_wallet_id"`
	TransactionType     string    `gorm:"type:varchar(100)" json:"transaction_type"`
	TransactionCategory int64     `gorm:"not null" json:"transaction_category_id"`
	TransactionDetail   string    `gorm:"type:text" json:"transaction_detail"`
	TransactionPrice    int64     `gorm:"not null" json:"transaction_price"`
	TransactionDate     time.Time `gorm:"not null" json:"transaction_date"`
}
