package response

type AuthResponse struct {
	Token      string `json:"token"`
	UserStatus string `json:"user_status"`
}
