package repository

import (
	"personal-blogging-api/model/domain"
	"time"

	"gorm.io/gorm"
)

type BlogRepositoryImpl struct {
	DB *gorm.DB
}

func NewBlogRepositoryImpl(db *gorm.DB) BlogRepository {
	return &BlogRepositoryImpl{DB: db}
}

func (b *BlogRepositoryImpl) Create(entity *domain.Blog) error {
	return b.DB.Create(entity).Error
}

func (b *BlogRepositoryImpl) Update(entity *domain.Blog) error {
	return b.DB.Save(entity).Error

}

func (b *BlogRepositoryImpl) Delete(id any) error {
	return b.DB.Delete(&domain.Blog{}, "id = ?", id).Error
}

func (b *BlogRepositoryImpl) FindById(id any, entity *domain.Blog) error {
	return b.DB.Where("id = ?", id).Take(entity).Error
}

func (b *BlogRepositoryImpl) FindAll() ([]domain.Blog, error) {
	var blogs []domain.Blog
	if err := b.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (b *BlogRepositoryImpl) FindBetweenTime(o1 time.Time, o2 time.Time) ([]domain.Blog, error) {
	var blogs []domain.Blog
	if err := b.DB.Where("created_at BETWEEN ? AND ?", o1, o2).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}
