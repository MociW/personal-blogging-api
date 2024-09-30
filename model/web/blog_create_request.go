package web

type BlogCreateRequest struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
