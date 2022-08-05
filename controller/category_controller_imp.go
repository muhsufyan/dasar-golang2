package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/model/web"
	"github.com/muhsufyan/dasar-golang2/service"
)

type CategoryControllerImp struct {
	// injek dg interface service
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImp) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// baca json, karena request.Body mrpkn io.Read maka kita bisa stream datanya memakai decoder
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	// decode data
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// update sama sprti create diatas bedanya kita perlu tahu dulu id berapa yg diupdate
	// baca json, karena request.Body mrpkn io.Read maka kita bisa stream datanya memakai decoder
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	// decode data
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)
	// get id yg ingin diupdate melalui params.ByName() yg return string jd kita perlu convert jd int
	categoryId := params.ByName("categoryId")
	// convert id from str to int
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get id yg ingin didelete melalui params.ByName() yg return string jd kita perlu convert jd int
	categoryId := params.ByName("categoryId")
	// convert id from str to int
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get id yg ingin didelete melalui params.ByName() yg return string jd kita perlu convert jd int
	categoryId := params.ByName("categoryId")
	// convert id from str to int
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponses := controller.CategoryService.FindAll(request.Context())
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoriesResponses,
	}
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
