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
	UserTglLahir  time.Time `json:"user_tgl_lahir"`
	UserAddress   string    `gorm:"type:text" json:"user_address"`
	UserCreatedAt time.Time `json:"user_created_at"`
	UserUpdatedAt time.Time `json:"user_updated_at"`
}
