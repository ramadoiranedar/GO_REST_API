package main

import (
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/ramadoiranedar/go_restapi/app"
	"github.com/ramadoiranedar/go_restapi/controller"
	"github.com/ramadoiranedar/go_restapi/exception"
	"github.com/ramadoiranedar/go_restapi/helper"
	"github.com/ramadoiranedar/go_restapi/middleware"
	"github.com/ramadoiranedar/go_restapi/repository"
	"github.com/ramadoiranedar/go_restapi/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categroyService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categroyService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
