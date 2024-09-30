package helper

import (
	"personal-blogging-api/model/domain"
	"personal-blogging-api/model/web"
)

func CreateResponse(blog domain.Blog) web.BlogResponse {
	response := web.BlogResponse{
		ID:      blog.ID,
		Author:  blog.Author,
		Title:   blog.Title,
		Content: blog.Content,
	}
	return response
}

func CreateResponses(blog []domain.Blog) []web.BlogResponse {
	var result []web.BlogResponse
	for _, data := range blog {
		result = append(result, CreateResponse(data))
	}
	return result
}
