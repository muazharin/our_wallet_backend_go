package request

type CategoryDeleteRequest struct {
	CategoryID int64 `form:"category_id" json:"category_id" binding:"required"`
}
