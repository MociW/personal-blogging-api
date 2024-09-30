package service

import (
	"personal-blogging-api/model/domain"
)

type BlogService interface {
	CreateBlog(entity *domain.Blog) error
	UpdateBlog(entity *domain.Blog) error
	DeleteBlog(id any) error
	FindBlogById(id any, entity *domain.Blog) error
	FindAllBlog() ([]domain.Blog, error)
	FindBetweenTimeBlog(o1 string) ([]domain.Blog, error)
}
