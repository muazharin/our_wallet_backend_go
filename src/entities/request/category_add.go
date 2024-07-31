package request

type CategoryAddRequest struct {
	CategoryTitle    string `form:"category_title" json:"category_title" binding:"required"`
	CategoryType     string `form:"category_type" json:"category_type" binding:"required"`
	CategoryWalletID string `form:"category_wallet_id" json:"category_wallet_id" binding:"required"`
}
