# dasar-golang2
DEPENDENCY INJECTION PADA PROJECT RESTAPI SBLMNYA (branch restapi-22-unittest)

Sblmnya di branch restapi-22-unittest kita melakukan DI scra manual, dpt dilihat di main.go dimana DI nya ada di code

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

tp hasil akhir yg dibthkan dari kode tsb sbnrnya adlh variabel server.
now kita akan jdkan server tsb as injector. untuk melakukan itu kita buat file baru di root directory injector.go


RESTAPI INI MENGIKKUTI TUTORIAL