package utils

import (
	"github.com/gin-gonic/gin"
)

// ResSuccess ...
func ResSuccess(data interface{}, msg string) (res *gin.H) {
	res = &gin.H{
		"code":    200,
		"data":    data,
		"message": msg,
	}
	return
}

// ResFaild ...
func ResFaild(msg string) (res *gin.H) {
	res = &gin.H{
		"code":    500,
		"data":    make(map[string]interface{}),
		"message": msg,
	}
	return
}

// ResOvertime ...
func ResOvertime() (res *gin.H) {
	res = &gin.H{
		"code":    400,
		"data":    make(map[string]interface{}),
		"message": "登录超时",
	}
	return
}
