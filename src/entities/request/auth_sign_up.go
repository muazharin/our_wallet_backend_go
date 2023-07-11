package request

type AuthSignUpRequest struct {
	UserName     string `form:"username" json:"username" binding:"required"`
	UserEmail    string `form:"email" json:"email" binding:"required"`
	UserPhone    string `form:"phone" json:"phone" binding:"required"`
	UserGender   string `form:"gender" json:"gender" binding:"required"`
	UserTglLahir string `form:"tgl_lahir" json:"tgl_lahir" binding:"required"`
	UserAddress  string `form:"address" json:"address" binding:"required"`
}
