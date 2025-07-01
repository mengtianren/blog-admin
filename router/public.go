package router

import (
	"blog-admin/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRouterInit(router *gin.Engine) {
	// 初始化路由
	r := router.Group("public")
	p := &controllers.PublicController{}
	r.POST("login", p.Login)
	r.POST("register", p.Register)
}
