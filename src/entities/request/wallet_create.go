package request

type WalletCreateReq struct {
	Name  string `form:"name" bson:"name" binding:"required"`
	Money int64  `form:"money" bson:"money" binding:"required"`
	Color string `form:"color" bson:"color" binding:"required"`
}
