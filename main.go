package main

import (
	"net/http"

	"github.com/ramadoiranedar/go_restapi/middleware"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadoiranedar/go_restapi/app"
	"github.com/ramadoiranedar/go_restapi/controller"
	"github.com/ramadoiranedar/go_restapi/helper"
	"github.com/ramadoiranedar/go_restapi/repository"
	"github.com/ramadoiranedar/go_restapi/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository(db)
	categroyService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categroyService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
