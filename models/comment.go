package models

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	Desc   string  `gorm:"desc;not null;comment:评论内容" json:"desc"`
	UserId uint    `gorm:"user_id;not null;comment:用户id" json:"user_id"`
	User   *User   `gorm:"foreignKey:UserId;references:ID" json:"user"` // 属于一个用户
	BlogId uint    `gorm:"blog_id;not null;comment:博客id" json:"blog_id"`
	Blog   *Blog   `gorm:"foreignKey:BlogId;references:ID" json:"blog"` // 属于一个博客
}
