package request

type AuthSignInRequest struct{
	UserName     	string `form:"username" json:"username" binding:"required"`
	UserPassword    string `form:"password" json:"password" binding:"required"`
}