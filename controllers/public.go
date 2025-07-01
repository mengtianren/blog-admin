package controllers

import (
	"blog-admin/core"
	"blog-admin/global"
	"blog-admin/services"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" required:"true" binding:"required,min=6,max=12"`
}

type PublicController struct {
}

func (p *PublicController) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.ResError(c, 400, "参数错误")
		return
	}
	s := &services.UserService{}
	// core.ResSuccess(c, 400, &req)

	user, err := s.GetUser(req.Phone, req.Password)
	if err != nil {
		core.ResError(c, 400, "用户名或密码错误")
		return
	}

	token, err := core.GenerateToken(user.ID, user.Phone)
	if err != nil || token == "" {
		core.ResError(c, 400, "登录失败")
		return
	}

	core.ResSuccess(c, gin.H{
		"token": token,
		"id":    user.ID,
	})

}

func (p *PublicController) Register(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.ResError(c, 400, "参数错误")
		return
	}
	s := &services.UserService{}
	global.Log.Infof("req: %v", req)
	if err := s.Register(req.Phone, req.Password); err != nil {
		core.ResError(c, 400, err.Error())
		return
	}

	core.ResSuccess(c, true)
}
