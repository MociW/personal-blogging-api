package repository

import (
	"personal-blogging-api/model/domain"
	"time"
)

type BlogRepository interface {
	Create(entity *domain.Blog) error
	Update(entity *domain.Blog) error
	Delete(id any) error
	FindById(id any, entity *domain.Blog) error
	FindAll() ([]domain.Blog, error)
	FindBetweenTime(o1 time.Time, o2 time.Time) ([]domain.Blog, error)
}
