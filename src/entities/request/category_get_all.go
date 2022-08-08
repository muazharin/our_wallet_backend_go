package request

type CategoryGetAllRequest struct {
	CategoryType     string `json:"category_type"`
	CategoryWalletId string `json:"category_wallet_id"`
	CategoryUserId   string `json:"category_user_id"`
	CategoryPage     int64  `json:"category_page"`
}
