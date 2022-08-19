package request

import "mime/multipart"

type TransCreateReq struct {
	TransWalletId   int64                   `form:"trans_wallet_id" bson:"trans_wallet_id" binding:"required"`
	TransType       string                  `form:"trans_type" bson:"trans_type" binding:"required"`
	TransCategoryId int64                   `form:"trans_category_id" bson:"trans_category_id" binding:"required"`
	TransDetail     string                  `form:"trans_detail" bson:"trans_detail" binding:"required"`
	TransPrice      int64                   `form:"trans_price" bson:"trans_price" binding:"required"`
	TransFile       []*multipart.FileHeader `form:"trans_file" bson:"trans_file" binding:""`
}
