package request

type OwGetUserReq struct {
	Page     int64  `form:"page" bson:"page" binding:"required"`
	WalletId int64  `form:"wallet_id" bson:"wallet_id" binding:"required"`
	Keyword  string `form:"keyword" bson:"keyword" binding:""`
}
