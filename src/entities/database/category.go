package database

type Category struct {
	CategoryID     int64  `gorm:"primary_key:auto_increment" json:"category_id"`
	CategoryUserID int64  `gorm:"not null" json:"-"`
	CategoryTitle  string `gorm:"type:varchar(100)" json:"category_title"`
	User           Users  `gorm:"foreignkey:CategoryUserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category_user_id"`
}
