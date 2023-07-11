package request

type WalletUpdateReq struct {
	WalletId int64  `form:"wallet_id" json:"wallet_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Money    int64  `form:"money" json:"money" binding:"required"`
	Color    string `form:"color" json:"color" binding:"required"`
}
