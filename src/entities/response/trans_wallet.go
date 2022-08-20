package response

type TransWallet struct {
	TransWalletID    int64  `json:"trans_wallet_id"`
	TransWalletName  string `json:"trans_wallet_name"`
	TransWalletColor string `json:"trans_wallet_color"`
}
