package database

type TransactionFile struct {
	TfID            int64       `gorm:"primary_key:auto_increment" json:"tf_id"`
	TfTramsactionID int64       `gorm:"not null" json:"-"`
	TfFile          string      `gorm:"type:varchar(100)" json:"tfFile"`
	Transaction     Transaction `gorm:"foreignkey:TfTramsactionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"tf_transaction_id"`
}
