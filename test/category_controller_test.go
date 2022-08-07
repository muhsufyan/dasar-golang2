package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/muhsufyan/dasar-golang2/app"
	"github.com/muhsufyan/dasar-golang2/controller"
	"github.com/muhsufyan/dasar-golang2/middleware"
	"github.com/muhsufyan/dasar-golang2/repository"
	"github.com/muhsufyan/dasar-golang2/service"
	"github.com/stretchr/testify/assert"
)

// sbnrnya fungsi ini sama sprti di main.go hanya beda dbnya saja
func setupRouterTest() http.Handler {
	// injeksikan dependensi yg diperlukan
	// 1a. injeksi dependensi koneksi ke db melalui db pooling (dependensi yg kita buat)
	db := app.TestDB()
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

// Buat skenario test
func TestCreateCategorySuccess(t *testing.T) {
	// setup router
	router := setupRouterTest()

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
func TestCreateCategoryFailed(t *testing.T) {

}
func TestUpdateCategorySuccess(t *testing.T) {

}
func TestUpdateCategoryFailed(t *testing.T) {

}
func TestGetCategorySuccess(t *testing.T) {

}
func TestGetCategoryFailed(t *testing.T) {

}
func TestDeleteCategorySuccess(t *testing.T) {

}
func TestDeleteCategoryFailed(t *testing.T) {

}
func TestGetListCategorySuccess(t *testing.T) {

}
func TestUnauthorized(t *testing.T) {

}
