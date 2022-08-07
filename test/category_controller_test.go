package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/middleware"
	"github.com/muhsufyan/dasar-golang2/model/domain"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
	"github.com/stretchr/testify/assert"
)

// sbnrnya fungsi ini sama sprti di main.go hanya beda dbnya saja
func setupRouterTest(db *sql.DB) http.Handler {
	// 1b. injeksi dependensi validasi data input
	validate := validator.New()
	// 2. injeksi dependensi repository (operasi query)
	categoryRepository := repository.NewCategoryRepository()
	//3. injeksi dependensi service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//4. injeksi dependensi controller
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

// agar data clear(data terdahulu dihapus untuk testing). ex create category pd test sblmnya akan dihapus menggunakan func ini
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

// Buat skenario test
func TestCreateCategorySuccess(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)
	// setup router
	router := setupRouterTest(db)

	// buat request body
	requestBody := strings.NewReader(`{"name":"dummy"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	// cek response
	assert.Equal(t, 200, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "dummy", responseBody["data"].(map[string]interface{})["name"])
}
func TestCreateCategoryFailedEmptyInput(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)
	// setup router
	router := setupRouterTest(db)

	// buat request body
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	// cek response
	assert.Equal(t, 400, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
func TestUpdateCategorySuccess(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// create dulu new data
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test",
	})
	tx.Commit()

	// setup router
	router := setupRouterTest(db)

	// buat request body
	requestBody := strings.NewReader(`{"name" : "dummy data testing"}`)
	fmt.Println("==================================")
	fmt.Println(strconv.Itoa(category.Id))
	fmt.Println("==================================")
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 200, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "dummy data testing", responseBody["data"].(map[string]interface{})["name"])
}
func TestUpdateCategoryFailedEmptyInput(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// create dulu new data
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test",
	})
	tx.Commit()

	// setup router
	router := setupRouterTest(db)

	// buat request body
	requestBody := strings.NewReader(`{"name" : ""}`)
	fmt.Println("==================================")
	fmt.Println(strconv.Itoa(category.Id))
	fmt.Println("==================================")
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	// cek response
	assert.Equal(t, 400, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
func TestGetCategoryByIdSuccess(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// create dulu new data
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test",
	})
	tx.Commit()

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 200, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}
func TestGetCategoryByIdFailedIdNotFound(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category/500", nil)
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 404, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}
func TestDeleteCategoryByIdSuccess(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// create dulu new data
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test",
	})
	tx.Commit()

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 200, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
}
func TestDeleteCategoryByIdFailedIdNotFound(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/category/500", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 404, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// convert float 64 jd int
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}
func TestGetListCategorySuccess(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// create dulu new data
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test first",
	})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "dummy test second",
	})
	tx.Commit()

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-KEY", "123ABC")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()
	fmt.Println(response.StatusCode)
	// cek response
	assert.Equal(t, 200, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// cek dulu datanya
	fmt.Println(responseBody)

	// convert float 64 jd int
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])

	// slice untuk menampung data response. convert jd slice map string
	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])

	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}
func TestUnauthorized(t *testing.T) {
	db := app.TestDB()
	// clear old data from test db
	truncateCategory(db)

	// setup router
	router := setupRouterTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-KEY", "WRONG")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// get response
	response := recorder.Result()

	// cek response
	assert.Equal(t, 401, response.StatusCode)

	// cek body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// cek dulu datanya
	fmt.Println(responseBody)

	// convert float 64 jd int
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
