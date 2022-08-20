package response

type TransUser struct {
	TransUserID    int64  `json:"trans_user_id"`
	TransUserName  string `json:"trans_user_name"`
	TransUserEmail string `json:"trans_user_email"`
	TransUserPhone string `json:"trans_user_phone"`
	TransUserPhoto string `json:"trans_user_photo"`
}
