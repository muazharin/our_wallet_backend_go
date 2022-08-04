package request

type AuthCheckPhoneRequest struct {
	Phone string `form:"phone" bson:"phone" binding:"required"`
}
