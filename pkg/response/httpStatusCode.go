package response

// Nội bộ đặt mã lỗi
const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20002 // Ví dụ báo lỗi email
)

// message
var msg = map[int]string{ // map kiểu int và string
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email invalid",
}
