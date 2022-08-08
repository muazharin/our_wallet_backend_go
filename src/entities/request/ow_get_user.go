package request

type OwGetUserReq struct {
	Page     int64 `json:"page"`
	WalletId int64 `json:"wallet_id"`
}
