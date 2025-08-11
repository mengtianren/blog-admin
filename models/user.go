package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string  `gorm:"name;size:10;default:'不知名用户';comment:用户名" json:"name"`
	Phone    string  `gorm:"phone;unique;size:11;not null;comment:手机号" json:"phone"`
	Password string  `gorm:"password;not null;comment:密码" json:"password"`
	Blogs    []*Blog `gorm:"foreignKey:UserId"` // 一个用户有多个博客
	Roles    []*Role `gorm:"many2many:user_roles"`
	Comments []*Comment `gorm:"foreignKey:UserId;references:ID" json:"comments"` // 一个用户有多个评论

	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}

type Role struct {
	ID    uint    `gorm:"id;primaryKey;autoIncrement;comment:主键"`
	Name  string  `gorm:"name;length:10;not null;comment:角色名称"`
	Users []*User `gorm:"many2many:user_roles"`
}

func (r *Role) TableName() string {
	return "role"
}
