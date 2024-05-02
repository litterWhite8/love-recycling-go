package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Data    interface{} `json:"data"`    // 数据
	Message string      `json:"message"` // 信息
}

type ResponseNoData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		200,
		data,
		msg,
	})
}

func SuccessNoData(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ResponseNoData{
		200,
		msg,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

func TokenFail(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		401,
		nil,
		"用户未登录",
	})
}
