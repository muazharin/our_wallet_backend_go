package database

type Category struct {
	CategoryID       int64  `gorm:"primary_key:auto_increment" json:"category_id"`
	CategoryWalletID int64  `gorm:"not null" json:"category_wallet_id"`
	CategoryTitle    string `gorm:"type:varchar(100)" json:"category_title"`
	CategoryType     string `gorm:"type:varchar(100)" json:"category_type"`
	CategoryIsActive bool   `gorm:"not null" json:"category_is_active"`
}
