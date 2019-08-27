package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/pkg/errno"
	"net/http"
)

// 返回固定格式 Data可以是 map、int、string、struct、array 等 Go 语言变量类型
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
