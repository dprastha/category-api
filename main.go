package main

import (
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/middleware"
	"net/http"
)

func main() {
	db := app.NewDB()
	router := app.SetupRouter(db)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
