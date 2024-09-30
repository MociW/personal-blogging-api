package service

import (
	"personal-blogging-api/model/domain"
	"personal-blogging-api/repository"
	"time"
)

type BlogServiceImpl struct {
	BlogRepository repository.BlogRepository
}

func NewBlogServiceImpl(blogRepository repository.BlogRepository) BlogService {
	return &BlogServiceImpl{blogRepository}
}

func (b *BlogServiceImpl) CreateBlog(entity *domain.Blog) error {
	return b.BlogRepository.Create(entity)
}

func (b *BlogServiceImpl) UpdateBlog(entity *domain.Blog) error {
	return b.BlogRepository.Update(entity)
}

func (b *BlogServiceImpl) DeleteBlog(id any) error {
	return b.BlogRepository.Delete(id)
}

func (b *BlogServiceImpl) FindBlogById(id any, entity *domain.Blog) error {
	return b.BlogRepository.FindById(id, entity)

}

func (b *BlogServiceImpl) FindAllBlog() ([]domain.Blog, error) {
	return b.BlogRepository.FindAll()
}

var (
	week    = time.Now().AddDate(0, 0, -7)
	week_2  = time.Now().AddDate(0, 0, -14)
	month   = time.Now().AddDate(0, -1, 0)
	month_2 = time.Now().AddDate(0, -2, 0)
	month_3 = time.Now().AddDate(0, -3, 0)
	month_6 = time.Now().AddDate(0, -6, 0)
	year    = time.Now().AddDate(-1, 0, 0)
)

func (b *BlogServiceImpl) FindBetweenTimeBlog(o1 string) ([]domain.Blog, error) {
	var date time.Time
	switch o1 {
	case "week":
		date = week
	case "2week":
		date = week_2
	case "month":
		date = month
	case "2month":
		date = month_2
	case "3month":
		date = month_3
	case "6month":
		date = month_6
	case "year":
		date = year
	}

	return b.BlogRepository.FindBetweenTime(date, time.Now())
}
