package model

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	UserID  uint
	UUID    string
	Title   string
	Content string
	Price   string
	Status  uint8 // 0:草稿, 1:已发布, 2:已下架
	Peony   string
}
