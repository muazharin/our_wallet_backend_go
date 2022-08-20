package database

type TransactionIsSeen struct {
	TransactionIsSeenID int64 `gorm:"primary_key:auto_increment" json:"transaction_is_seen_id"`
	TransactionUserID   int64 `gorm:"not null" json:"transaction_user_id"`
	TransactionID       int64 `gorm:"not null" json:"transaction_id"`
}
