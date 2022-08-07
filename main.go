package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/exception"
	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/middleware"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
)

func main() {
	// injeksikan dependensi yg diperlukan
	// 1a. injeksi dependensi koneksi ke db melalui db pooling (dependensi yg kita buat)
	db := app.NewDB()
	// 1b. injeksi dependensi validasi data input
	validate := validator.New()
	// 2. injeksi dependensi repository (operasi query)
	categoryRepository := repository.NewCategoryRepository()
	//3. injeksi dependensi service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//4. injeksi dependensi controller
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)
	// change panic error to error handler
	router.PanicHandler = exception.ErrorHandler
	// http server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
