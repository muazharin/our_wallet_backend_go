package request

type AuthSignOutRequest struct {
	UserFirebaseToken string `form:"firebase_token" json:"firebase_token" binding:"required"`
}
