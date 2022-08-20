package request

type TransByWalletIdReq struct {
	TransWalletId int64 `form:"trans_wallet_id" bson:"trans_wallet_id" binding:"required"`
	Page          int64 `form:"page" bson:"page" binding:"required"`
}
