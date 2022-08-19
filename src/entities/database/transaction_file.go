package database

type TransactionFile struct {
	TfID            int64  `gorm:"primary_key:auto_increment" json:"tf_id"`
	TfTransactionID int64  `gorm:"not null" json:"tf_transaction_id"`
	TfFile          string `gorm:"type:varchar(100)" json:"tfFile"`
}
