package request

type WalletCreateReq struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Money int64  `form:"money" json:"money" binding:"required"`
	Color string `form:"color" json:"color" binding:"required"`
}
