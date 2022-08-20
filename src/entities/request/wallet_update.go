package request

type WalletUpdateReq struct {
	WalletId int64  `form:"wallet_id" bson:"wallet_id" binding:"required"`
	Name     string `form:"name" bson:"name" binding:"required"`
	Money    int64  `form:"money" bson:"money" binding:"required"`
	Color    string `form:"color" bson:"color" binding:"required"`
}
