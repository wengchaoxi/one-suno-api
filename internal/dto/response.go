package dto

import "github.com/gin-gonic/gin"

const (
	STATUS_SUCCESS      = 0
	STATUS_WAITING      = -1
	STATUS_TIMEOUT      = -2
	STATUS_UNAUTHORIZED = -3
	STATUS_SERVER_ERROR = -4
)

type Status int

var (
	STATUS_MSG_MAP = map[Status]string{
		STATUS_SUCCESS:      "Success",
		STATUS_WAITING:      "Waiting, The Server is busy now, please try again later.",
		STATUS_TIMEOUT:      "Timeout, Please try again later.",
		STATUS_UNAUTHORIZED: "Unauthorized, Please provide a valid api key.",
		STATUS_SERVER_ERROR: "Server error, Please try again later.",
	}
)

type Response struct {
	Code    Status `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

func SendResponse(ctx *gin.Context, status Status, data any) {
	ctx.JSON(200, &Response{
		Code:    status,
		Message: STATUS_MSG_MAP[status],
		Data:    data,
	})
}
