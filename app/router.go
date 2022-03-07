package app

import (
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/exception"
	"belajar-golang-rest-api/middleware"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/service"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	// Categories
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	// products
	router.GET("/api/v1/products", productController.FindAll)
	router.GET("/api/v1/products/:productId", productController.FindById)
	router.POST("/api/v1/products", productController.Create)
	router.PUT("/api/v1/products/:productId", productController.Update)
	router.DELETE("/api/v1/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := NewRouter(categoryController, productController)

	return middleware.NewAuthMiddleware(router)
}
