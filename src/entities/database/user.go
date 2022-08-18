package database

import (
	"time"
)

type Users struct {
	UserID        int64     `gorm:"primary_key:auto_increment" json:"user_id"`
	UserName      string    `gorm:"type:varchar(100)" json:"user_name"`
	UserPassword  string    `gorm:"type:varchar(100)" json:"user_password"`
	UserEmail     string    `gorm:"type:varchar(100)" json:"user_email"`
	UserPhone     string    `gorm:"type:varchar(100)" json:"user_phone"`
	UserPhoto     string    `gorm:"type:varchar(100)" json:"user_photo"`
	UserGender    string    `gorm:"type:varchar(6)" json:"user_gender"`
	UserTglLahir  time.Time `gorm:"not null" json:"user_tgl_lahir"`
	UserAddress   string    `gorm:"type:text" json:"user_address"`
	UserStatus    string    `gorm:"type:text" json:"user_status"`
	UserCreatedAt time.Time `gorm:"not null" json:"user_created_at"`
	UserUpdatedAt time.Time `gorm:"not null" json:"user_updated_at"`
}
