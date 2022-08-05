package controller

import (
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
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	// response ke user
	helper.Write2ResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)
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
	// response ke user
	helper.Write2ResponseBody(writer, webResponse)
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
	// response ke user
	helper.Write2ResponseBody(writer, webResponse)
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
	// response ke user
	helper.Write2ResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponses := controller.CategoryService.FindAll(request.Context())
	// data model untuk response ke user
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoriesResponses,
	}
	// response ke user
	helper.Write2ResponseBody(writer, webResponse)
}
