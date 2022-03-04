package main

import (
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/middleware"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
