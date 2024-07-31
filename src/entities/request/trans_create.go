package request

import "mime/multipart"

type TransCreateReq struct {
	TransWalletId   int64                   `form:"trans_wallet_id" json:"trans_wallet_id" binding:"required"`
	TransType       string                  `form:"trans_type" json:"trans_type" binding:"required"`
	TransCategoryId int64                   `form:"trans_category_id" json:"trans_category_id" binding:"required"`
	TransDetail     string                  `form:"trans_detail" json:"trans_detail" binding:"required"`
	TransPrice      int64                   `form:"trans_price" json:"trans_price" binding:"required"`
	TransFile       []*multipart.FileHeader `form:"trans_file" json:"trans_file" binding:""`
}
