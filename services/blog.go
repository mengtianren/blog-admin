package services

import (
	"blog-admin/global"
	"blog-admin/models"
	"errors"
)

// 博客详情响应
type blogRes struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	UserPhone string `json:"user_phone"`
}
type BlogService struct {
	*models.Blog
}

func (b *BlogService) GetById(id uint) (*blogRes, error) {
	var blog models.Blog
	err := global.DB.Preload("User").First(&blog, id).Error
	if err != nil {
		return nil, err
	}
	res := &blogRes{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		UserID:    blog.User.ID,
		UserPhone: blog.User.Phone,
	}

	return res, nil
}
func (b *BlogService) Create(userId uint, title, content string) error {
	blog := &models.Blog{
		UserId:  userId,
		Title:   title,
		Content: content,
	}
	err := global.DB.Create(&blog).Error
	if err != nil {
		return err
	}
	return nil
}
func (b *BlogService) Update(userId, id uint, title, content string) error {
	blog := &models.Blog{
		Title:   title,
		Content: content,
	}
	var bl models.Blog
	if err := global.DB.Where("id = ?", id).First(&bl).Error; err != nil {
		return err
	}
	if bl.UserId != userId {
		return errors.New("只能修改自己发布的文章")
	}

	err := global.DB.Model(&models.Blog{}).Where("id = ?", id).Updates(&blog).Error
	if err != nil {
		return err
	}
	return nil
}
