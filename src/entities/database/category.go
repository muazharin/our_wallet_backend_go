package database

type Category struct {
	CategoryID       int64  `gorm:"primary_key:auto_increment" json:"category_id"`
	CategoryUserID   int64  `gorm:"not null" json:"category_user_id"`
	CategoryWalletID int64  `gorm:"not null" json:"category_wallet_id"`
	CategoryTitle    string `gorm:"type:varchar(100)" json:"category_title"`
	CategoryType     string `gorm:"type:varchar(100)" json:"category_type"`
}
