package domain

import "gorm.io/gorm"

type Blog struct {
	*gorm.Model
	Author  string `gorm:"column:author"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (b *Blog) TableName() string {
	return "blogs"
}
