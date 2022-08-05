package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// kita memakai package http router, bisa juga memakai gin, mux, dll

type CategoryController interface {
	// htttprouter perlu params jd arg 3 perlu ditambah
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
