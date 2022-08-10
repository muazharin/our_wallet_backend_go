package request

type AuthSignUpRequest struct {
	UserName     string `form:"username" bson:"username" binding:"required"`
	UserEmail    string `form:"email" bson:"email" binding:"required"`
	UserPhone    string `form:"phone" bson:"phone" binding:"required"`
	UserGender   string `form:"gender" bson:"gender" binding:"required"`
	UserTglLahir string `form:"tgl_lahir" bson:"tgl_lahir" binding:"required"`
	UserAddress  string `form:"address" bson:"address" binding:"required"`
}
