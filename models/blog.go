package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	Title   string `gorm:"title;size:32;not null;comment:标题 " json:"title"`
	Content string `gorm:"content;size:255;not null;comment:内容;" json:"content"`
	User    *User  `gorm:"foreignKey:UserId;references:ID" json:"user"` // belongs To
	UserId  uint   `gorm:"user_id;not null;comment:用户id" json:"user_id"`
	*gorm.Model
}

func (b *Blog) TableName() string {
	return "blog"
}
