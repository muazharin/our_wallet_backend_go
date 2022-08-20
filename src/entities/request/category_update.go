package request

type CategoryUpdateRequest struct {
	CategoryID    int64  `form:"category_id" bson:"category_id" binding:"required"`
	CategoryTitle string `form:"category_title" bson:"category_title" binding:"required"`
}
