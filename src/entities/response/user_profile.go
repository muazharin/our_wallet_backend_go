package response

type UserProfileRes struct {
	UserID        int64  `json:"user_id"`
	UserName      string `json:"user_name"`
	UserEmail     string `json:"user_email"`
	UserPhone     string `json:"user_phone"`
	UserPhoto     string `json:"user_photo"`
	UserGender    string `json:"user_gender"`
	UserTglLahir  string `json:"user_tgl_lahir"`
	UserAddress   string `json:"user_address"`
	UserCreatedAt string `json:"user_created_at"`
	UserUpdatedAt string `json:"user_updated_at"`
}
