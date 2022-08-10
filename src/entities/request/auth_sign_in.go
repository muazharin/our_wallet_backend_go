package request

type AuthSignInRequest struct {
	UserName     string `form:"username" bson:"username" binding:"required"`
	UserPassword string `form:"password" bson:"password" binding:"required"`
}
