package database

type Notification struct {
	NotificationID      int64  `gorm:"primary_key:auto_increment" json:"notification_id"`
	NotificationUserID  int64  `gorm:"not null" json:"notification_user_id"`
	NotificationMessage string `gorm:"type:text" json:"notification_message"`
	NotificationRoute   string `gorm:"type:varchar(100)" json:"notification_route"`
	NotificationIsRead  bool   `gorm:"not null" json:"notification_is_read"`
}
