package controller

import (
	"fmt"
	"personal-blogging-api/helper"
	"personal-blogging-api/model/domain"
	"personal-blogging-api/service"

	"github.com/gofiber/fiber/v2"
)

type BlogControllerImpl struct {
	BlogService service.BlogService
}

func NewControllerImpl(blogService service.BlogService) BlogController {
	return &BlogControllerImpl{BlogService: blogService}
}

func (b *BlogControllerImpl) CreateBlog(c *fiber.Ctx) error {
	var entity domain.Blog
	if err := c.BodyParser(&entity); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := b.BlogService.CreateBlog(&entity); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := helper.CreateResponse(entity)
	return c.Status(200).JSON(response)
}

func (b *BlogControllerImpl) UpdateBlog(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	// Fetch the existing blog entry
	entity := new(domain.Blog)
	if err := b.BlogService.FindBlogById(id, entity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Parse the update request body
	if err := c.BodyParser(entity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update the blog using the service
	if err := b.BlogService.UpdateBlog(entity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create and return the response
	return c.Status(fiber.StatusOK).JSON(helper.CreateResponse(*entity))
}

func (b *BlogControllerImpl) DeleteBlog(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if err := b.BlogService.DeleteBlog(id); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Successfully deleted User")
}

func (b *BlogControllerImpl) FindBlogById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	entity := new(domain.Blog)
	err := b.BlogService.FindBlogById(id, entity)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := helper.CreateResponse(*entity)
	return c.Status(200).JSON(response)

}

func (b *BlogControllerImpl) FindAllBlog(c *fiber.Ctx) error {

	blogs, err := b.BlogService.FindAllBlog()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := helper.CreateResponses(blogs)
	return c.Status(200).JSON(response)
}

func (b *BlogControllerImpl) FindBetweenTimeBlog(c *fiber.Ctx) error {
	period := c.Query("period")
	fmt.Println("Period: ", period)

	blogs, err := b.BlogService.FindBetweenTimeBlog(period)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := helper.CreateResponses(blogs)
	return c.Status(200).JSON(response)
}
