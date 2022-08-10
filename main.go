package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/middleware"
)

// provider server
func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {
	// NOW CALL SERVER DARI HASIL GENERATE WIRE
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

/*fungsi main dg manual dependency
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

	router := app.NewRouter(categoryController)

	// 5. injek dependensi middleware
	authMiddleware := middleware.NewAuthMiddleware(router)
	// 6. injeksi dependensi http server
	server := NewServer(authMiddleware)
*/
