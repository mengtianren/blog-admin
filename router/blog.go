package router

import (
	"blog-admin/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRouterInit(r *gin.Engine) {
	// 初始化路由
	router := r.Group("blog")
	c := &controllers.BlogController{}

	router.GET("", c.GetBlog)
	router.POST("", c.PostBlog)
	router.PUT("", c.PutBlog)
}
