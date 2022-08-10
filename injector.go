// go:build wireinject
// +build wireinject
package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
)

// buat func injector
func InitializedServer() *http.Server {
	wire.Build(
		// karena NewDB returnnya struct jd tdk perlu binding interface
		app.NewDB,
		// ini juga struct
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
	)
	return nil
}
