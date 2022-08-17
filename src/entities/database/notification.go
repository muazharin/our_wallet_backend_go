package database

type Notification struct {
	NotificationID         int64  `gorm:"primary_key:auto_increment" json:"notification_id"`
	NotificationPusherID   int64  `gorm:"not null" json:"notification_pusher_id"`
	NotificationReceiverID int64  `gorm:"not null" json:"notification_receiver_id"`
	NotificationMessage    string `gorm:"type:text" json:"notification_message"`
	NotificationRoute      string `gorm:"type:varchar(100)" json:"notification_route"`
	NotificationIsRead     bool   `gorm:"not null" json:"notification_is_read"`
}
