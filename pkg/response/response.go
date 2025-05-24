package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessRepsonse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
