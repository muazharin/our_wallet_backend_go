package request

type TransByIdReq struct {
	TransId int64 `form:"trans_id" bson:"trans_id" binding:"required"`
}
