package request

type CategoryDeleteRequest struct {
	CategoryID int64 `form:"category_id" bson:"category_id" binding:"required"`
}
