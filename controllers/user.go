package controllers

import (
	"blog-admin/core"
	"blog-admin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*services.UserService
}

func (u *UserController) UserInfo(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		core.ResError(c, http.StatusUnauthorized, "未登录")
		return
	}
	user, err := u.GetUserById(userId.(uint))
	if err != nil {
		core.ResError(c, http.StatusInternalServerError, "获取用户信息失败")
		return
	}

	core.ResSuccess(c, gin.H{
		"userId": user,
	})
}
