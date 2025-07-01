package core

import (
	"blog-admin/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResError(c *gin.Context, code int, msg string) {
	global.Log.Errorf("请求错误: %s", msg)
	c.JSON(http.StatusNotImplemented, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func ResSuccess[T interface{}](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": data,
	})
}
