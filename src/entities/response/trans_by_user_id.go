package response

type TransByUserIdRes struct {
	TransID       int64       `json:"trans_id"`
	TransType     string      `json:"trans_type"`
	TransCategory string      `json:"trans_category"`
	TransDetail   string      `json:"trans_detail"`
	TransPrice    int64       `json:"trans_price"`
	TransDate     string      `json:"trans_date"`
	TransWallet   TransWallet `json:"trans_wallet"`
}
