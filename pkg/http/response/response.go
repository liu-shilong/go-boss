package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 响应
type Response struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// 成功
func Success(c *gin.Context, data interface{}) {
	code := http.StatusOK
	c.JSON(code, Response{
		Code: code,
		Msg:  "success",
		Data: data,
	})
}

// 失败
func Fail(c *gin.Context) {
	code := http.StatusBadRequest
	c.JSON(code, Response{
		Code: code,
		Msg:  "fail",
		Data: nil,
	})
}
