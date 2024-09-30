package web

type BlogResponse struct {
	ID      uint   `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
