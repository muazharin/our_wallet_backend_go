package request

type OwConfirmInvitation struct {
	ConfirmWalletId int64 `form:"confirm_wallet_id" json:"confirm_wallet_id" binding:"required"`
	ConfirmReply    bool  `form:"confirm_reply" json:"confirm_reply" binding:""`
}
