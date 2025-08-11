package router

import (
	"blog-admin/controllers"
	"blog-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func BlogRouterInit(r *gin.Engine) {
	// 初始化路由
	router := r.Group("blog")
	c := &controllers.BlogController{}
	router.POST("list", c.GetBlogList)
	router.GET("", c.GetBlog)
	router.POST("", middlewares.JwtMiddleware(), c.PostBlog)
	router.PUT("", middlewares.JwtMiddleware(), c.PutBlog)
	router.POST("comment", middlewares.JwtMiddleware(), c.PostComment)
}
