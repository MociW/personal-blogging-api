package controller

import (
	"github.com/gofiber/fiber/v2"
)

type BlogController interface {
	CreateBlog(c *fiber.Ctx) error
	UpdateBlog(c *fiber.Ctx) error
	DeleteBlog(c *fiber.Ctx) error
	FindBlogById(c *fiber.Ctx) error
	FindAllBlog(c *fiber.Ctx) error
	FindBetweenTimeBlog(c *fiber.Ctx) error
}
