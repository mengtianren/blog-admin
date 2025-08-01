package router

import (
	"blog-admin/controllers"
	"blog-admin/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.Engine) {
	// 初始化路由
	r := router.Group("user")
	r.Use(middlewares.JwtMiddleware())
	c := controllers.UserController{}
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "user"})
	})
	r.GET("info", c.UserInfo)
	r.PUT("password", c.UpdateP)
}
