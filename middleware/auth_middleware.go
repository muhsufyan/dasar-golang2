package middleware

import (
	"net/http"

	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/model/web"
)

// buat struct dg contract handler server
type AuthMiddleware struct {
	// next action after auth (ex next to page after handler auth is ok)
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}
func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// ex api key nya 123ABC
	if "123ABC" == request.Header.Get("X-API-KEY") {
		// OK, lanjutkan handler selanjutnya (controller)
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		// response ke user
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.Write2ResponseBody(writer, webResponse)
	}
}
