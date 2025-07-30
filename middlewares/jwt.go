package middlewares

import (
	"blog-admin/core"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// token 验证中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if strings.HasPrefix(c.Request.URL.Path, "/public") {
		// 	c.Next()
		// 	return
		// }

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			core.ResError(c, http.StatusUnauthorized, "token不存在")
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := core.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户权限有误"})
			c.Abort()
			return
		}

		fmt.Printf("token信息是：%+v \n", claims)

		c.Set("userId", claims.UserId)
		c.Set("Phone", claims.Phone)
		c.Next()
	}
}
