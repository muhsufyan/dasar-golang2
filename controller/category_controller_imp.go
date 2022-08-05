package controller

import (
	"net/http"

	"github.com/go-delve/delve/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImp struct {
	// injek dg interface service
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImp) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// baca json
}
func (controller *CategoryControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
func (controller *CategoryControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
func (controller *CategoryControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
func (controller *CategoryControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
