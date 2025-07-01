package router

import (
	"blog-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	// 初始化路由
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.LoggerMiddleware(), middlewares.JwtMiddleware())
	{
		BlogRouterInit(r)
		UserRouterInit(r)
		PublicRouterInit(r)
	}

	return r
}
