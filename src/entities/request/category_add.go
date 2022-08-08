package request

type CategoryAddRequest struct {
	CategoryTitle    string `form:"category_title" bson:"category_title" binding:"required"`
	CategoryType     string `form:"category_type" bson:"category_type" binding:"required"`
	CategoryWalletID string `form:"category_wallet_id" bson:"category_wallet_id" binding:"required"`
}
