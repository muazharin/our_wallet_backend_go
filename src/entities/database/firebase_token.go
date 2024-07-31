package database

type FirebaseToken struct {
	FirebaseTokenID     int64  `gorm:"primary_key:auto_increment" json:"firebase_token_id"`
	FirebaseTokenUserID int64  `gorm:"not null" json:"firebase_user_id"`
	FirebaseTokenString string `gorm:"type:text" json:"firebase_token_string"`
}
