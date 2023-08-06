package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int
	Message string
	Data    interface{}
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, response{Code: http.StatusOK, Message: message, Data: data})
}

func ReturnData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{Code: http.StatusOK, Message: "成功", Data: data})
}

func SimpleSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, response{Code: http.StatusOK, Message: "执行成功", Data: nil})
}

func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, response{Code: http.StatusInternalServerError, Message: "服务内部错误"})
}

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, response{Code: http.StatusBadRequest, Message: "请求数据错误"})
}
