package main

import (
	"dedenurr/golang-restful-api/app"
	"dedenurr/golang-restful-api/controller"
	"dedenurr/golang-restful-api/helper"
	"dedenurr/golang-restful-api/middleware"
	"dedenurr/golang-restful-api/repository"
	"dedenurr/golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate :=validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService :=  service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	

}