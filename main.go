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

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(categoryController, productController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
