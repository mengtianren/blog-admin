package core

import "gorm.io/gorm"

type Page struct {
	Page  int                    `json:"page" form:"page" default:"1"`
	Size  int                    `json:"size" form:"size" default:"10"`
	Param map[string]interface{} `json:"param" form:"param"`
}

type PageResponse[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
	Size  int   `json:"size"`
	Page  int   `json:"page"`
}

func Paginate[T any](db *gorm.DB, page Page, out *[]T) (int64, error) {
	if page.Page <= 0 {
		page.Page = 1
	}
	if page.Size <= 0 {
		page.Size = 10
	}

	var total int64
	db.Count(&total)

	offset := (page.Page - 1) * page.Size
	err := db.Limit(page.Size).Offset(offset).Find(out).Error
	return total, err
}
