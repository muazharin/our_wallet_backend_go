package request

type CategoryUpdateRequest struct {
	CategoryID    int64  `form:"category_id" json:"category_id" binding:"required"`
	CategoryTitle string `form:"category_title" json:"category_title" binding:"required"`
}
