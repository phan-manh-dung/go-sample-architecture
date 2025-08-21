package response

// Nội bộ đặt mã lỗi
const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20002 // Ví dụ báo lỗi email
	ErrInvalidToken     = 20003 // Invalid token
)

// message
var msg = map[int]string{ // map kiểu int và string
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email invalid",
	ErrInvalidToken:     "Invalid token",
}
