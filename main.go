package main

import (
	"personal-blogging-api/app"
	"personal-blogging-api/controller"
	"personal-blogging-api/repository"
	"personal-blogging-api/service"
)

func main() {
	db := app.NewDB()
	blogRepository := repository.NewBlogRepositoryImpl(db)
	blogService := service.NewBlogServiceImpl(blogRepository)
	blogController := controller.NewControllerImpl(blogService)
	router := app.NewRouter(blogController)

	router.Listen(":3000")
}
