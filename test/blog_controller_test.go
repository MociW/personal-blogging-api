package test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"personal-blogging-api/app"
	"personal-blogging-api/controller"
	"personal-blogging-api/repository"
	"personal-blogging-api/service"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var conn = fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error: " + err.Error())
	},
})

func SetUp() controller.BlogController {
	db := app.NewDB()
	blogRepository := repository.NewBlogRepositoryImpl(db)
	blogService := service.NewBlogServiceImpl(blogRepository)
	blogController := controller.NewControllerImpl(blogService)
	return blogController
}

func TestCreateBlog(t *testing.T) {

	conn.Post("/v1/blog", SetUp().CreateBlog)

	body := strings.NewReader(`{"author":"wina","title":"Lorem Ipsum","content":"Lorem ipsum dolor sit amet, consectetur adipiscing elit."}`)
	request := httptest.NewRequest("POST", "/v1/blog", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}

func TestUpdateBlog(t *testing.T) {

	conn.Put("/v1/blog/:id", SetUp().UpdateBlog)

	body := strings.NewReader(`{"author":"wina oktalia","title":"Lorem Ipsum","content":"Lorem ipsum dolor sit amet, consectetur adipiscing elit."}`)
	request := httptest.NewRequest("PUT", "/v1/blog/6", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}

func TestDeleteBlog(t *testing.T) {
	conn.Delete("/v1/blog/:id", SetUp().DeleteBlog)

	request := httptest.NewRequest("DELETE", "/v1/blog/5", nil)
	request.Header.Set("Content-Type", "application/json")
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}

func TestGetById(t *testing.T) {
	conn.Get("/v1/blog/:id", SetUp().FindBlogById)

	request := httptest.NewRequest("GET", "/v1/blog/4", nil)
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}

func TestGetAll(t *testing.T) {
	conn.Get("/v1/blog", SetUp().FindAllBlog)

	request := httptest.NewRequest("GET", "/v1/blog", nil)
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}

func TestGetBetweenTime(t *testing.T) {
	conn.Get("/v1/blog", SetUp().FindAllBlog)

	request := httptest.NewRequest("GET", "/v1/blog?period=week", nil)
	response, err := conn.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}
