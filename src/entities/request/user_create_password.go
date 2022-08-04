package request

type UserCreatePasswordRequest struct {
	Password string `form:"password" bson:"password" binding:"required"`
}
