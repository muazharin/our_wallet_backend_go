package request

type IsReadNotifReq struct {
	NotifId int64 `form:"notif_id" json:"notif_id" binding:"required"`
}
