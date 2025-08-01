package models

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	Desc   string `json:"desc"`
	UserId uint   `json:"user_id"`
	User   User   `json:"user"`
	BlogId uint   `json:"blog_id"`
	Blog   Blog   `json:"blog"`
}
