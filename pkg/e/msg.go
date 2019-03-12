package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	NO_ERROR  = 0
	HAS_ERROR = 1
)

func resSuccess(data interface{}) map[string]interface{} {
	return gin.H{
		"code":  http.StatusOK,
		"error": NO_ERROR,
		"data":  data,
	}
}

func resWarning(data interface{}) map[string]interface{} {
	return gin.H{
		"code":  http.StatusOK,
		"error": HAS_ERROR,
		"data":  data,
	}
}

func resOutput(code int, data interface{}) map[string]interface{} {
	return gin.H{
		"code":  code,
		"error": HAS_ERROR,
		"data":  data,
	}
}

//成功输出
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resSuccess(data))
}

//警告输出
func Warning(c *gin.Context, err error) {
	c.JSON(http.StatusOK, resWarning(err.Error()))
}

//自定义输出
func Output(c *gin.Context, code int, data interface{}) {
	c.JSON(code, resOutput(code, data))
}
