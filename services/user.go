package services

import (
	"blog-admin/core"
	"blog-admin/global"
	"blog-admin/models"
	"errors"

	"gorm.io/gorm"
)

type Info struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Roles []uint `json:"roles"`
}
type UserService struct {
	*models.User
}

func (u *UserService) Register(phone string, password string) error {
	user := &models.User{}
	err := global.DB.Where("phone = ?", phone).First(&user).Error
	if err == nil {
		return errors.New("用户已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("注册失败")
	}
	if p, err := core.HashPassword(password); err != nil {
		return errors.New("注册失败")
	} else {
		user.Password = p
	}

	user.Phone = phone
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil

}
func (u *UserService) GetUserById(id uint) (*Info, error) {
	user := &models.User{}

	err := global.DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	var roles = []uint{}
	for _, v := range user.Roles {
		roles = append(roles, v.ID)
	}

	info := &Info{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Roles: roles,
	}

	return info, nil
}

func (u *UserService) GetUser(phone string, password string) (*Info, error) {
	user := &UserService{}

	err := global.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	success := core.CheckPasswordHash(password, user.Password)
	if !success {
		return nil, errors.New("账号或密码错误")
	}

	info := &Info{
		ID:    user.ID,
		Phone: user.Phone,
		Name:  user.Name,
	}
	return info, nil

}
func (u *UserService) UpdateInfo(id uint, phone string) error {
	var user = &models.User{
		Phone: phone,
	}
	err := global.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *UserService) UpdatePassword(id uint, old, password string) error {
	var user models.User
	if err := global.DB.First(&user, id).Error; err != nil {
		return err
	}
	if !core.CheckPasswordHash(old, user.Password) {
		return errors.New("旧密码错误")
	}
	p, err1 := core.HashPassword(password)
	if err1 != nil {
		return err1
	}
	global.DB.Model(&user).Update("password", p)
	return nil
}
