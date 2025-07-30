package services

import (
	"blog-admin/core"
	"blog-admin/global"
	"blog-admin/models"
	"errors"
	"time"
)

// 博客详情响应
type blogRes struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	ShowNum  int    `json:"show_num"`
}
type BlogService struct {
	*models.Blog
}
type BlogList struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" `
	Content   string    `json:"content" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `json:"user_id"`
}

func (b *BlogService) List(page core.Page) (*core.PageResponse[BlogList], error) {
	db := global.DB.Model(&models.Blog{})
	var blogs []BlogList
	total, err := core.Paginate(db, page, &blogs)
	if err != nil {
		return nil, err
	}
	if page.Size == 0 {
		page.Size = 10
	}
	if page.Page == 0 {
		page.Page = 1
	}

	return &core.PageResponse[BlogList]{
		List:  blogs,
		Total: total,
		Size:  page.Size,
		Page:  page.Page,
	}, nil

}

func (b *BlogService) GetById(id uint) (*blogRes, error) {
	var blog models.Blog

	db := global.DB.Model(&models.Blog{}).Where("id = ?", id)

	err := db.Preload("User").First(&blog).Error
	if err != nil {
		return nil, err
	}
	db.Updates(&models.Blog{ShowNum: blog.ShowNum + 1})

	res := &blogRes{
		ID:       blog.ID,
		Title:    blog.Title,
		Content:  blog.Content,
		UserID:   blog.User.ID,
		UserName: blog.User.Name,
		ShowNum:  blog.ShowNum,
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
