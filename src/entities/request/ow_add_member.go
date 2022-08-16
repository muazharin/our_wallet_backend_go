package request

type OwAddMemberReq struct {
	OwWalletId int64 `form:"ow_wallet_id" json:"ow_wallet_id" binding:"required"`
	OwMemberId int64 `form:"ow_member_id" json:"ow_member_id" binding:"required"`
}
