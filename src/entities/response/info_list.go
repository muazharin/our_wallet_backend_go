package response

type InfoResponse struct {
	InfoPage      int64 `json:"info_page"`
	InfoLimit     int64 `json:"info_limit"`
	InfoTotalData int64 `json:"info_total_data"`
	InfoTotalPage int64 `json:"info_total_page"`
}
