package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // mã lỗi
	Message string      `json:"message"` // message
	Data    interface{} `json:"data"`    // dữ liệu được return
}

// Success
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    ErrCodeSuccess,
		Message: msg[code],
		Data:    data,
	})
}

// Error
func ErrorResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    ErrCodeSuccess,
		Message: msg[code],
		Data:    nil,
	})
}
