package app

import (
	"personal-blogging-api/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(c controller.BlogController) *fiber.App {
	app := fiber.New()
	app.Get("/v1/blog/search", c.FindBetweenTimeBlog)
	app.Get("/v1/blog", c.FindAllBlog)
	app.Post("/v1/blog", c.CreateBlog)
	app.Put("/v1/blog/:id", c.UpdateBlog)
	app.Delete("/v1/blog/:id", c.DeleteBlog)
	app.Get("/v1/blog/:id", c.FindBlogById)
	return app
}
