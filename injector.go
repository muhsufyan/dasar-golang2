//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/middleware"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
)

// provider set
// provider repository kita ubah dulu returnnya (kalau tetap interface itu aman/bisa langsung tanpa binding), jd kita bljr binding
// provider service juga kita ubah returnnya agar belajar binding sama sprti repository
// provider controller juga kita ubah returnnya agar belajar binding sama sprti repository & service
var categorySet = wire.NewSet(
	// lakukan binding untuk repository
	repository.NewCategoryRepository,
	// kalau ada yg buth CategoryRepository (interface) maka kirimkan CategoryRepositoryImpl (struct)
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	// lakukan binding untuk service
	service.NewCategoryService,
	// kalau ada yg buth CategoryService (interface) maka kirimkan CategoryServiceImpl (struct)
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	// lakukan binding untuk controller
	controller.NewCategoryController,
	// kalau ada yg buth CategoryController (interface) maka kirimkan CategoryControllerImpl (struct)
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImp)),
)

// buat func injector
func InitializedServer() *http.Server {
	wire.Build(
		// karena NewDB returnnya struct jd tdk perlu binding interface
		app.NewDB,
		// ini juga struct
		validator.New,
		categorySet,
		//provider func NewRouter bth (paramnya) CategoryController
		app.NewRouter,
		// binding http.Handler, kenapa perlu melakukan ini ? karena NewRouter itu returnnya *httprouter.Router
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		// auth middleware perlu (paramnya) http.Handler yg berupa interface
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
