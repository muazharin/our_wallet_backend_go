package request

type AuthSignInRequest struct {
	UserName          string `form:"username" json:"username" binding:"required"`
	UserPassword      string `form:"password" json:"password" binding:"required"`
	UserFirebaseToken string `form:"firebase_token" json:"firebase_token" binding:"required"`
}
