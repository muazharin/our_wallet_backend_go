package request

type TransByUserIdReq struct {
	TransUserId int64 `form:"trans_user_id" bson:"trans_user_id" binding:"required"`
	Page        int64 `form:"page" bson:"page" binding:"required"`
}
