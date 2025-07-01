package controllers

import (
	"blog-admin/core"
	"blog-admin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type baseReq struct {
	ID      uint   `json:"id" `
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type BlogController struct {
	*services.BlogService
}

func (b *BlogController) GetBlog(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		core.ResError(c, 400, "id 不存在")
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		core.ResError(c, 500, "id 必须是数字")
		return
	}

	data, err := b.GetById(uint(id))
	if err != nil {
		core.ResError(c, 400, "文章不存在")
		return
	}

	core.ResSuccess(c, data)
}
func (b *BlogController) PostBlog(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		core.ResError(c, http.StatusUnauthorized, "未登录")
		return
	}
	var req baseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.ResError(c, 400, "请检查参数")
		return
	}

	if err := b.Create(userId.(uint), req.Title, req.Content); err != nil {
		core.ResError(c, 400, err.Error())
		return
	}

	core.ResSuccess(c, true)

}

func (b *BlogController) PutBlog(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		core.ResError(c, http.StatusUnauthorized, "未登录")
		return
	}

	var req baseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.ResError(c, 400, "请检查参数")
		return
	}
	if req.ID == 0 {
		core.ResError(c, 400, "id 不存在")
		return
	}

	if err := b.Update(userId.(uint), req.ID, req.Title, req.Content); err != nil {
		core.ResError(c, 400, err.Error())
		return
	}
	core.ResSuccess(c, true)
}
