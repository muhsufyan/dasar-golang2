// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/middleware"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

// buat func injector
func InitializedServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImp := controller.NewCategoryController(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImp)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

// provider set
// provider repository kita ubah dulu returnnya (kalau tetap interface itu aman/bisa langsung tanpa binding), jd kita bljr binding
// provider service juga kita ubah returnnya agar belajar binding sama sprti repository
// provider controller juga kita ubah returnnya agar belajar binding sama sprti repository & service
var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImp)))
