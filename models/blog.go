package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	Title   string `gorm:"title;size:32;not null;comment:标题 " json:"title"`
	Content string `gorm:"content;size:255;not null;comment:内容;" json:"content"`
	Type    string `gorm:"type:enum('类型1','类型2');default:'类型1'" comment:"类型"`
	User    *User  `gorm:"foreignKey:UserId;references:ID" json:"user"` // belongs To
	ShowNum int    `gorm:"show_num;default:0" comment:"阅读量"`
	UserId  uint   `gorm:"user_id;not null;comment:用户id" json:"user_id"`
	*gorm.Model
}

func (b *Blog) TableName() string {
	return "blog"
}
