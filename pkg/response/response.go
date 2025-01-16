package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success_response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})

}
