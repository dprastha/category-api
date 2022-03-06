package app

import (
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/exception"

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
