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

func (u *UserController) UserRoles(c *gin.Context) {

}
func (u *UserController) UpdateP(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		core.ResError(c, http.StatusUnauthorized, "未登录")
		return
	}

	var update struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}
	if c.ShouldBind(&update) != nil {
		core.ResError(c, http.StatusBadRequest, "参数绑定失败")
		return
	}

	err := u.UpdatePassword(userId.(uint), update.OldPassword, update.Password)

	if err != nil {
		core.ResError(c, 500, err.Error())
		return
	}
	core.ResSuccess(c, true)
}
