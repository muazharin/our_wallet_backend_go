package response

type TransByWalletIdRes struct {
	TransID       int64        `json:"trans_id"`
	TransType     string       `json:"trans_type"`
	TransCategory string       `json:"trans_category"`
	TransDetail   string       `json:"trans_detail"`
	TransPrice    int64        `json:"trans_price"`
	TransDate     string       `json:"trans_date"`
	TransIsSeen   bool         `json:"trans_is_seen"`
	TransUser     TransUser    `json:"trans_user"`
	TransWallet   *TransWallet `json:"trans_wallet"`
}
